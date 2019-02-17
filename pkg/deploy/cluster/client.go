package cluster

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

func (c *Cluster) createEndpoint(endpoint map[string]string) error {
	endpoint["parent_type"] = defaultResource
	endpoint["display_name"] = endpoint["name"]
	endpoint["prefix"] = endpoint["name"]
	endpoint["name"] = fmt.Sprintf("%s-%s", endpoint["name"], uuid.NewV4().String())
	endpointData := map[string]map[string]string{"endpoint": endpoint}
	c.log.Infof("Creating endpoint: %s, %s", endpoint["display_name"], endpoint["public_url"])
	var endpointResponse map[string]interface{}
	resURI := fmt.Sprintf("%ss", defaultEndpointResPath)
	_, err := c.APIServer.Create(context.Background(), resURI, &endpointData, &endpointResponse)
	return err
}

func (c *Cluster) getDefaultCredential() (user, password, keypair string, err error) {
	var credList map[string][]interface{}
	resURI := fmt.Sprintf("%ss", defaultCredentialResPath)
	c.log.Infof("Reading credential: %s", resURI)
	_, err = c.APIServer.Read(context.Background(), resURI, &credList)
	if err != nil {
		return "", "", "", err
	}
	for _, rawCred := range credList[defaultCredentialRes+"s"] {
		cred := models.InterfaceToCredential(rawCred)
		if cred.Name == "default-credential" {
			for _, keypairRef := range cred.KeypairRefs {
				k, err := c.getResource(defaultKeypairResPath, keypairRef.UUID)
				if err != nil {
					return "", "", "", err
				}
				keypair := models.InterfaceToKeypair(k)
				if keypair.Name == "default-keypair" {
					return cred.SSHUser, cred.SSHPassword, keypair.SSHPublicKey, nil
				}
			}
			return cred.SSHUser, cred.SSHPassword, "", nil
		}
	}
	return "", "", "", nil

}

func (c *Cluster) getEndpoints(parentUUIDs []string) (endpointIDs []string, err error) {
	values := url.Values{
		baseservices.ParentUUIDsKey: parentUUIDs,
		baseservices.ParentTypeKey:  []string{defaultResource},
	}
	var endpointList map[string][]interface{}
	resURI := fmt.Sprintf("%ss?%s", defaultEndpointResPath, values.Encode())
	c.log.Infof("Reading endpoints: %s", resURI)
	_, err = c.APIServer.Read(context.Background(), resURI, &endpointList)
	if err != nil {
		return nil, err
	}
	for _, rawEndpoint := range endpointList[defaultEndpointRes+"s"] {
		endpointID := rawEndpoint.(map[string]interface{})["uuid"].(string) // nolint: errcheck
		endpointIDs = append(endpointIDs, endpointID)
	}
	return endpointIDs, nil
}

func (c *Cluster) deleteEndpoint(endpointUUID string) error {
	var output map[string]interface{}
	resURI := fmt.Sprintf("%s/%s", defaultEndpointResPath, endpointUUID)
	c.log.Infof("Deleting endpoint: %s", resURI)
	//TODO(nati) fixed context
	_, err := c.APIServer.Delete(context.Background(), resURI, &output)
	return err
}

func (c *Cluster) getResource(resPath string, resID string) (map[string]interface{}, error) {
	var rawResInfo map[string]interface{}
	resURI := fmt.Sprintf("%s/%s", resPath, resID)
	c.log.Infof("Reading: %s", resURI)
	_, err := c.APIServer.Read(context.Background(), resURI, &rawResInfo)
	if err != nil {
		return nil, err
	}
	res := strings.TrimLeft(resPath, "/")
	data, ok := rawResInfo[res].(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid resource type")
	}
	return data, nil
}

func (c *Cluster) getNode(nodeID string, m map[string]bool, d DataStore) error {
	if _, ok := m[nodeID]; !ok {
		m[nodeID] = true
		n, err := c.getResource(defaultNodeResPath, nodeID)
		if err != nil {
			return err
		}
		ni := models.InterfaceToNode(n)
		d.addNode(ni)
		for _, credRef := range ni.CredentialRefs {
			// stop iteration after one credential ref
			return c.getCredential(credRef.UUID, m, d)
		}
	}
	return nil
}

func (c *Cluster) getKeypair(keypairID string, m map[string]bool, d DataStore) error {
	if _, ok := m[keypairID]; !ok {
		m[keypairID] = true
		k, err := c.getResource(defaultKeypairResPath, keypairID)
		if err != nil {
			return err
		}
		d.addKeypair(models.InterfaceToKeypair(k))
		return nil
	}
	return nil
}

func (c *Cluster) getCredential(credentialID string, m map[string]bool, d DataStore) error {
	if _, ok := m[credentialID]; !ok {
		m[credentialID] = true
		ci, err := c.getResource(defaultCredentialResPath, credentialID)
		if err != nil {
			return err
		}
		cred := models.InterfaceToCredential(ci)
		d.addCredential(cred)
		for _, keypairRef := range cred.KeypairRefs {
			// stop iteration after one keypair ref
			return c.getKeypair(keypairRef.UUID, m, d)
		}
	}
	return nil
}

func (c *Cluster) getClusterDetails(clusterID string) (*Data, error) {
	// get contrail cluster information
	clusterData := &Data{Reader: c.APIServer}
	if err := clusterData.updateClusterDetails(clusterID, c); err != nil {
		return nil, err
	}

	// get all referred openstack cluster information
	for _, openstackClusterRef := range clusterData.clusterInfo.OpenstackClusterRefs {
		openstakData := &OpenstackData{Reader: c.APIServer}
		if err := openstakData.updateClusterDetails(
			openstackClusterRef.UUID, c); err != nil {
			return nil, err
		}
		clusterData.openstackClusterData = append(
			clusterData.openstackClusterData, openstakData)
	}

	// get all referred kubernetes cluster information
	for _, kubernetesClusterRef := range clusterData.clusterInfo.KubernetesClusterRefs {
		k8sData := &KubernetesData{Reader: c.APIServer}
		if err := k8sData.updateClusterDetails(
			kubernetesClusterRef.UUID, c); err != nil {
			return nil, err
		}
		clusterData.kubernetesClusterData = append(
			clusterData.kubernetesClusterData, k8sData)
	}

	// get all referred vCenter information
	for _, vcenterRef := range clusterData.clusterInfo.VCenterRefs {
		vCenterData := &VCenterData{Reader: c.APIServer}
		if err := vCenterData.updateClusterDetails(
			vcenterRef.UUID, c); err != nil {
			return nil, err
		}
		clusterData.vcenterData = append(
			clusterData.vcenterData, vCenterData)
	}

	// get all referred appformix cluster information
	for _, appformixClusterRef := range clusterData.clusterInfo.AppformixClusterRefs {
		appformixData := &AppformixData{Reader: c.APIServer}
		if err := appformixData.updateClusterDetails(
			appformixClusterRef.UUID, c); err != nil {
			return nil, err
		}
		clusterData.appformixClusterData = append(
			clusterData.appformixClusterData, appformixData)
	}

	return clusterData, nil
}
