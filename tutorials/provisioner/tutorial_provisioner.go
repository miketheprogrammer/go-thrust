package tutorial

import "github.com/cloudspace/go-thrust/lib/spawn"

/*
Default implementation of Provisioner
*/
type TutorialProvisioner struct{}

func NewTutorialProvisioner() TutorialProvisioner {
	return TutorialProvisioner{}
}

func (tp TutorialProvisioner) Provision() error {
	spawn.SetBaseDirectory("") // Means use the users home directory
	return spawn.Bootstrap()

}
