package cloud

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	tf "github.com/hashicorp/terraform/terraform"
)

func readStateFile(tfStateFile string) (*tf.State, error) {

	if _, err := os.Stat(tfStateFile); err == nil {
		state, err := ioutil.ReadFile(tfStateFile)
		if err != nil {
			return nil, err
		}
		if len(state) == 0 {
			return nil, fmt.Errorf("tfState does not contain any data")
		}

		stateBuf := bytes.NewBuffer(state)
		tfState, err := tf.ReadState(stateBuf)
		if err != nil {
			return nil, err
		}
		return tfState, nil
	}
	return nil, fmt.Errorf("tf state file: %s does not exist", tfStateFile)
}

func getIPFromTFState(tfState *tf.State, outputKey string) (string, error) {

	mState := tfState.RootModule()
	output, ok := mState.Outputs[outputKey]
	if !ok {
		return "", fmt.Errorf("output key %s doesn't exist in tfState", outputKey)
	}

	return output.Value.(string), nil
}
