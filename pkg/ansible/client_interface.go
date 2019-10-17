package ansible

// Player runs ansibleplaybooks
type Player interface {
	Play()
}

const (
	playbookCmd    = "ansible-playbook"
	filePermRWOnly = 0600
)
