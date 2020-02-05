package cloud

import (
	"context"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	cloudmock "github.com/Juniper/contrail/pkg/cloud/mock"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
)

const (
	computeRole = "compute"
	gatewayRole = "gateway"

	privatePortName = "private"
	hostName        = "hostName"
	publicIP        = "1.1.1.1"
	privateIP       = "2.2.2.2"
	uuid            = "uuid"
)

type updateIPDetailsTestSuite struct {
	suite.Suite
	dummyContext   context.Context
	mockCtrl       *gomock.Controller
	httpClientMock *servicesmock.MockService
	apiMocked      apiServer
	tfStateMock    *cloudmock.MockTerraformState
}

func (s *updateIPDetailsTestSuite) SetupTest() {
	s.dummyContext = context.Background()
	s.mockCtrl = gomock.NewController(s.T())
	s.httpClientMock = servicesmock.NewMockService(s.mockCtrl)
	s.apiMocked = apiServer{s.httpClientMock, s.dummyContext}
	s.tfStateMock = cloudmock.NewMockTerraformState(s.mockCtrl)
}

func (s *updateIPDetailsTestSuite) TeardownTest() {
	s.mockCtrl.Finish()
}

func (s *updateIPDetailsTestSuite) TestUpdateNotGatewayInstancePublicIP() {
	instance := newInstance(hostName, computeRole, uuid, s.apiMocked)

	s.tfStateMock.EXPECT().GetPrivateIP(hostName).Return(privateIP, nil)
	s.httpClientMock.EXPECT().UpdateNode(gomock.Any(), nodeIPUpdateMatcher{hostName, privateIP})

	s.NoError(updateIPDetails(s.dummyContext, []*instanceData{instance}, s.tfStateMock))
}

func (s *updateIPDetailsTestSuite) TestUpdateGatewayInstancePublicIPAndCreatePrivatePortWhenPortIsNotPresent() {

	instance := newInstance(hostName, gatewayRole, uuid, s.apiMocked)

	s.tfStateMock.EXPECT().GetPrivateIP(hostName).Return(privateIP, nil)
	s.tfStateMock.EXPECT().GetPublicIP(hostName).Return(publicIP, nil)

	s.httpClientMock.EXPECT().CreatePort(gomock.Any(), createPrivatePortMatcher{uuid, privateIP}).Return(
		&services.CreatePortResponse{
			Port: &models.Port{
				Name:       privatePortName,
				ParentType: models.KindNode,
				ParentUUID: uuid,
				IPAddress:  privateIP}}, nil)
	s.httpClientMock.EXPECT().UpdateNode(gomock.Any(), updateNodePrivatePortMatcher{uuid, hostName, privateIP})
	s.httpClientMock.EXPECT().UpdateNode(gomock.Any(), nodeIPUpdateMatcher{hostName, publicIP})

	s.NoError(updateIPDetails(s.dummyContext, []*instanceData{instance}, s.tfStateMock))
}

func (s *updateIPDetailsTestSuite) TestUpdateGatewayInstancePublicIPAndUpdatePrivatePortIPWhenPrivatePortIsPresent() {
	otherIP := "0.0.0.0"

	instance := newInstance(hostName, gatewayRole, uuid, s.apiMocked)
	instance.info.Ports = []*models.Port{{
		Name:       privatePortName,
		ParentType: models.KindNode,
		ParentUUID: uuid,
		IPAddress:  otherIP}}

	s.tfStateMock.EXPECT().GetPrivateIP(hostName).Return(privateIP, nil)
	s.tfStateMock.EXPECT().GetPublicIP(hostName).Return(publicIP, nil)

	s.httpClientMock.EXPECT().UpdatePort(gomock.Any(), updatePrivatePortMatcher{uuid, privateIP}).Return(
		&services.UpdatePortResponse{
			Port: &models.Port{
				Name:       privatePortName,
				ParentType: models.KindNode,
				ParentUUID: uuid,
				IPAddress:  privateIP}}, nil)
	s.httpClientMock.EXPECT().UpdateNode(gomock.Any(), updateNodePrivatePortMatcher{uuid, hostName, privateIP})
	s.httpClientMock.EXPECT().UpdateNode(gomock.Any(), nodeIPUpdateMatcher{hostName, publicIP})

	s.NoError(updateIPDetails(s.dummyContext, []*instanceData{instance}, s.tfStateMock))
}

func (s *updateIPDetailsTestSuite) TestUpdateGatewayInstancePublicIPOnly() {
	instance := newInstance(hostName, gatewayRole, uuid, s.apiMocked)
	instance.info.Ports = []*models.Port{{
		Name:       privatePortName,
		ParentType: models.KindNode,
		ParentUUID: uuid,
		IPAddress:  privateIP}}

	s.tfStateMock.EXPECT().GetPrivateIP(hostName).Return(privateIP, nil)
	s.tfStateMock.EXPECT().GetPublicIP(hostName).Return(publicIP, nil)

	s.httpClientMock.EXPECT().UpdateNode(gomock.Any(), updateNodePrivatePortMatcher{uuid, hostName, privateIP})
	s.httpClientMock.EXPECT().UpdateNode(gomock.Any(), nodeIPUpdateMatcher{hostName, publicIP})

	s.NoError(updateIPDetails(s.dummyContext, []*instanceData{instance}, s.tfStateMock))
}

func newInstance(hostName string, role string, uuid string, api apiServer) *instanceData {
	return &instanceData{
		info: &models.Node{
			CloudInfo: &models.CloudInstanceInfo{
				Roles: []string{role},
			},
			Hostname: hostName,
			UUID:     uuid,
		},
		apiServer: api,
	}
}

func TestUpdateIPDetails(t *testing.T) {
	suite.Run(t, new(updateIPDetailsTestSuite))
}

type updateNodePrivatePortMatcher struct {
	expectedUUID     string
	expectedHostname string
	expectedIP       string
}

func (m updateNodePrivatePortMatcher) Matches(x interface{}) bool {
	request, ok := x.(*services.UpdateNodeRequest)
	return ok &&
		request.Node.Hostname == m.expectedHostname &&
		len(request.Node.Ports) > 0 &&
		request.Node.Ports[0].Name == privatePortName &&
		request.Node.Ports[0].ParentType == models.KindNode &&
		request.Node.Ports[0].ParentUUID == m.expectedUUID &&
		request.Node.Ports[0].IPAddress == m.expectedIP
}

func (m updateNodePrivatePortMatcher) String() string {
	return fmt.Sprintf(
		"is equal to UpdateNodeRequest with Node.Hostname: %v ,"+
			"Node.Port{Name: %v, ParentType: %v ,ParentUUID:%v, IPAddress: %v",
		m.expectedHostname, privatePortName, models.KindNode, m.expectedUUID, m.expectedIP)
}

type updatePrivatePortMatcher struct {
	expectedUUID string
	expectedIP   string
}

func (m updatePrivatePortMatcher) Matches(x interface{}) bool {
	request, ok := x.(*services.UpdatePortRequest)
	return ok &&
		request.Port.Name == privatePortName &&
		request.Port.ParentType == models.KindNode &&
		request.Port.ParentUUID == m.expectedUUID &&
		request.Port.IPAddress == m.expectedIP
}

func (m updatePrivatePortMatcher) String() string {
	return fmt.Sprintf("is equal to UpdateNodeRequest with Port{Name: %v, ParentType: %v ,ParentUUID:%v, IPAddress: %v",
		privatePortName, models.KindNode, m.expectedUUID, m.expectedIP)
}

type createPrivatePortMatcher struct {
	expectedUUID string
	expectedIP   string
}

func (m createPrivatePortMatcher) Matches(x interface{}) bool {
	request, ok := x.(*services.CreatePortRequest)
	return ok &&
		request.Port.Name == privatePortName &&
		request.Port.ParentType == models.KindNode &&
		request.Port.ParentUUID == m.expectedUUID &&
		request.Port.IPAddress == m.expectedIP
}

func (m createPrivatePortMatcher) String() string {
	return fmt.Sprintf("is equal to CreatePortRequest with Port{Name: %v, ParentType: %v ,ParentUUID:%v, IPAddress: %v",
		privatePortName, models.KindNode, m.expectedUUID, m.expectedIP)
}

type nodeIPUpdateMatcher struct {
	hostName   string
	expectedIP string
}

func (m nodeIPUpdateMatcher) Matches(x interface{}) bool {
	request, ok := x.(*services.UpdateNodeRequest)
	return ok && request.Node.IPAddress == m.expectedIP && request.Node.Hostname == m.hostName
}

func (m nodeIPUpdateMatcher) String() string {
	return fmt.Sprintf("is equal to UpdateNodeRequest with Node{HostName: %v, IPAddress: %v",
		m.hostName, m.expectedIP)
}
