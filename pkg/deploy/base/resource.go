package base

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/format"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

// ResourceManager to manage resources
type ResourceManager struct {
	APIServer *client.HTTP
	Log       *logrus.Entry
}

// NewResourceManager creates ResourceManager
func NewResourceManager(APIServer *client.HTTP, logFile string) *ResourceManager {
	return &ResourceManager{
		APIServer: APIServer,
		Log:       logutil.NewFileLogger("resource-manager", logFile),
	}
}

func (r *ResourceManager) createEndpoint(endpoint map[string]string) error {
	endpoint["parent_type"] = defaultResource
	endpoint["display_name"] = endpoint["name"]
	endpoint["prefix"] = endpoint["name"]
	endpoint["name"] = fmt.Sprintf("%s-%s", endpoint["name"], uuid.NewV4().String())
	endpointData := map[string]map[string]string{"endpoint": endpoint}
	r.Log.Infof("Creating endpoint: %s, %s", endpoint["display_name"], endpoint["public_url"])
	var endpointResponse map[string]interface{}
	resURI := fmt.Sprintf("%ss", defaultEndpointResPath)
	_, err := r.APIServer.Create(context.Background(), resURI, &endpointData, &endpointResponse)
	return err
}

func (r *ResourceManager) getDefaultCredential() (user, password, keypair string, err error) {
	var credList map[string][]interface{}
	resURI := fmt.Sprintf("%ss", defaultCredentialResPath)
	r.Log.Infof("Reading credential: %s", resURI)
	_, err = r.APIServer.Read(context.Background(), resURI, &credList)
	if err != nil {
		return "", "", "", err
	}
	for _, rawCred := range credList[defaultCredentialRes+"s"] {
		cred := models.InterfaceToCredential(rawCred)
		if cred.Name == "default-credential" {
			for _, keypairRef := range cred.KeypairRefs {
				k, err := r.getResource(defaultKeypairResPath, keypairRef.UUID)
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

func (r *ResourceManager) getEndpoints(parentUUIDs []string) (endpointIDs []string, err error) {
	values := url.Values{
		baseservices.ParentUUIDsKey: parentUUIDs,
		baseservices.ParentTypeKey:  []string{defaultResource},
		"prefix":                    format.GetKeys(portMap),
	}
	var endpointList map[string][]interface{}
	resURI := fmt.Sprintf("%ss?%s", defaultEndpointResPath, values.Encode())
	r.Log.Infof("Reading endpoints: %s", resURI)
	_, err = r.APIServer.Read(context.Background(), resURI, &endpointList)
	if err != nil {
		return nil, err
	}
	for _, rawEndpoint := range endpointList[defaultEndpointRes+"s"] {
		endpointID := rawEndpoint.(map[string]interface{})["uuid"].(string) // nolint: errcheck
		endpointIDs = append(endpointIDs, endpointID)
	}
	return endpointIDs, nil
}

func (r *ResourceManager) deleteEndpoint(endpointUUID string) error {
	var output map[string]interface{}
	resURI := fmt.Sprintf("%s/%s", defaultEndpointResPath, endpointUUID)
	r.Log.Infof("Deleting endpoint: %s", resURI)
	//TODO(nati) fixed context
	_, err := r.APIServer.Delete(context.Background(), resURI, &output)
	return err
}

func (r *ResourceManager) getResource(resPath string, resID string) (map[string]interface{}, error) {
	var rawResInfo map[string]interface{}
	resURI := fmt.Sprintf("%s/%s", resPath, resID)
	r.Log.Infof("Reading: %s", resURI)
	_, err := r.APIServer.Read(context.Background(), resURI, &rawResInfo)
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

func (r *ResourceManager) getNode(nodeID string, m map[string]bool, d DataStore) error {
	if _, ok := m[nodeID]; !ok {
		m[nodeID] = true
		n, err := r.getResource(defaultNodeResPath, nodeID)
		if err != nil {
			return err
		}
		ni := models.InterfaceToNode(n)
		d.addNode(ni)
		for _, credRef := range ni.CredentialRefs {
			// stop iteration after one credential ref
			return r.getCredential(credRef.UUID, m, d)
		}
	}
	return nil
}

func (r *ResourceManager) getKeypair(keypairID string, m map[string]bool, d DataStore) error {
	if _, ok := m[keypairID]; !ok {
		m[keypairID] = true
		k, err := r.getResource(defaultKeypairResPath, keypairID)
		if err != nil {
			return err
		}
		d.addKeypair(models.InterfaceToKeypair(k))
		return nil
	}
	return nil
}

func (r *ResourceManager) getCredential(credentialID string, m map[string]bool, d DataStore) error {
	if _, ok := m[credentialID]; !ok {
		m[credentialID] = true
		ci, err := r.getResource(defaultCredentialResPath, credentialID)
		if err != nil {
			return err
		}
		cred := models.InterfaceToCredential(ci)
		d.addCredential(cred)
		for _, keypairRef := range cred.KeypairRefs {
			// stop iteration after one keypair ref
			return r.getKeypair(keypairRef.UUID, m, d)
		}
	}
	return nil
}

// GetClusterDetails gets contrail cluster details
func (r *ResourceManager) GetClusterDetails(clusterID string) (*Data, error) { // nolint:gocyclo
	// get contrail cluster information
	clusterData := &Data{Reader: r.APIServer}
	if err := clusterData.updateClusterDetails(clusterID, r); err != nil {
		return nil, err
	}

	// get all referred openstack cluster information
	for _, openstackClusterRef := range clusterData.ClusterInfo.OpenstackClusterRefs {
		openstakData := &OpenstackData{Reader: r.APIServer}
		if err := openstakData.updateClusterDetails(
			openstackClusterRef.UUID, r); err != nil {
			return nil, err
		}
		clusterData.openstackClusterData = append(
			clusterData.openstackClusterData, openstakData)
	}

	// get all referred kubernetes cluster information
	for _, kubernetesClusterRef := range clusterData.ClusterInfo.KubernetesClusterRefs {
		k8sData := &KubernetesData{Reader: r.APIServer}
		if err := k8sData.updateClusterDetails(
			kubernetesClusterRef.UUID, r); err != nil {
			return nil, err
		}
		clusterData.kubernetesClusterData = append(
			clusterData.kubernetesClusterData, k8sData)
	}

	// get all referred vCenter information
	for _, vcenterRef := range clusterData.ClusterInfo.VCenterRefs {
		vCenterData := &VCenterData{Reader: r.APIServer}
		if err := vCenterData.updateClusterDetails(
			vcenterRef.UUID, r); err != nil {
			return nil, err
		}
		clusterData.vcenterData = append(
			clusterData.vcenterData, vCenterData)
	}

	// get all referred appformix cluster information
	for _, appformixClusterRef := range clusterData.ClusterInfo.AppformixClusterRefs {
		appformixData := &AppformixData{Reader: r.APIServer}
		if err := appformixData.updateClusterDetails(
			appformixClusterRef.UUID, r); err != nil {
			return nil, err
		}

		for _, appformixFlows := range appformixData.ClusterInfo.AppformixFlowss {
			xflowData := NewXflowData()
			if err := xflowData.updateClusterDetails(context.Background(), appformixFlows.UUID, r); err != nil {
				return nil, err
			}
			clusterData.xflowData = append(clusterData.xflowData, xflowData)
		}

		clusterData.appformixClusterData = append(
			clusterData.appformixClusterData, appformixData)
	}

	return clusterData, nil
}
