package cluster

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"context"

	uuid "github.com/satori/go.uuid"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

func (c *Cluster) createEndpoint(parentUUID, name, publicURL, privateURL string) error {
	endpoint := map[string]string{
		"parent_type":  defaultResource,
		"parent_uuid":  parentUUID,
		"name":         fmt.Sprintf("%s-%s", name, uuid.NewV4().String()),
		"display_name": name,
		"prefix":       name,
		"public_url":   publicURL,
		"private_url":  privateURL,
	}
	endpointData := map[string]map[string]string{"endpoint": endpoint}
	c.log.Infof("Creating endpoint: %s, %s", name, publicURL)
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
		services.ParentUUIDsKey: parentUUIDs,
		services.ParentTypeKey:  []string{defaultResource},
	}
	var endpointList map[string][]interface{}
	resURI := fmt.Sprintf("%ss?%s", defaultEndpointResPath, values.Encode())
	c.log.Infof("Reading endpoints: %s", resURI)
	_, err = c.APIServer.Read(context.Background(), resURI, &endpointList)
	if err != nil {
		return nil, err
	}
	for _, rawEndpoint := range endpointList[defaultEndpointRes+"s"] {
		endpointID := rawEndpoint.(map[string]interface{})["uuid"].(string)
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
	clusterData := &Data{}
	if err := clusterData.updateClusterDetails(clusterID, c); err != nil {
		return nil, err
	}

	// get all referred openstack cluster information
	for _, openstackClusterRef := range clusterData.clusterInfo.OpenstackClusterRefs {
		openstakData := &OpenstackData{}
		if err := openstakData.updateClusterDetails(
			openstackClusterRef.UUID, c); err != nil {
			return nil, err
		}
		clusterData.openstackClusterData = append(
			clusterData.openstackClusterData, openstakData)
	}

	// get all referred kubernetes cluster information
	for _, kubernetesClusterRef := range clusterData.clusterInfo.KubernetesClusterRefs {
		k8sData := &KubernetesData{}
		if err := k8sData.updateClusterDetails(
			kubernetesClusterRef.UUID, c); err != nil {
			return nil, err
		}
		clusterData.kubernetesClusterData = append(
			clusterData.kubernetesClusterData, k8sData)
	}

	return clusterData, nil
}
