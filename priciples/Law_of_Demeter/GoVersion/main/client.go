package main

import "LODGoversion"

func main() {
	agent := new(LODGoversion.Agent)
	agent.SetCompany(*LODGoversion.NewCompanyByName("希望传媒"))
	agent.SetStart(*LODGoversion.NewStartByName("周杰伦"))
	agent.SetFans(*LODGoversion.NewFansByName("小明"))

	agent.Meeting()
	agent.Business()
}