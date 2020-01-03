package cloud

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/terraform/terraform"
)

type terraformState interface {
	GetPublicIP(hostname string) (string, error)
	GetPrivateIP(hostname string) (string, error)
}

type cloudTfStateReader struct {
	cloudUUID string
	dryRun    bool
}

func (r cloudTfStateReader) Read() (terraformState, error) {
	if r.dryRun {
		return dryRunTfState{}, nil
	}
	tfState, err := readStateFile(GetTFStateFile(r.cloudUUID))
	if err != nil {
		return nil, err
	}
	return &cloudTfState{tfState: tfState}, nil
}

const dumyIP = "1.2.3.4"

type dryRunTfState struct {
}

func (s dryRunTfState) GetPublicIP(hostname string) (string, error) {
	return dumyIP, nil
}

func (s dryRunTfState) GetPrivateIP(hostname string) (string, error) {
	return dumyIP, nil
}

type cloudTfState struct {
	tfState *terraform.State
}

func (s *cloudTfState) GetPublicIP(hostname string) (string, error) {
	return s.getKeyValue(fmt.Sprintf("%s.public_ip", hostname))
}

func (s *cloudTfState) GetPrivateIP(hostname string) (string, error) {
	return s.getKeyValue(fmt.Sprintf("%s.private_ip", hostname))
}

func (s *cloudTfState) getKeyValue(outputKey string) (string, error) {
	mState := s.tfState.RootModule()
	output, ok := mState.Outputs[outputKey]
	if !ok {
		return "", fmt.Errorf("output key %s doesn't exist in tfState", outputKey)
	}

	return output.Value.(string), nil
}

func readStateFile(tfStateFile string) (*terraform.State, error) {

	if _, err := os.Stat(tfStateFile); err == nil {
		state, err := ioutil.ReadFile(tfStateFile)
		if err != nil {
			return nil, err
		}
		if len(state) == 0 {
			return nil, fmt.Errorf("tfState does not contain any data")
		}

		stateBuf := bytes.NewBuffer(state)
		tfState, err := terraform.ReadState(stateBuf)
		if err != nil {
			return nil, err
		}
		return tfState, nil
	}
	return nil, fmt.Errorf("terraform state file: %s does not exist", tfStateFile)
}
