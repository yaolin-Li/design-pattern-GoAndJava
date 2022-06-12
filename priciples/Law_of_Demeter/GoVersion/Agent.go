package LODGoversion

import "fmt"

type Agent struct{
	fans Fans
	start Start
	company Company
}

func (a *Agent) SetFans(fans Fans) {
	a.fans = fans
} 

func (a *Agent) SetStart(start Start) {
	a.start = start
}

func (a *Agent) SetCompany(company Company) {
	a.company = company
} 

func (a Agent) Meeting() {
	fmt.Println(a.start.getName()+"和粉丝"+a.fans.getName()+"见面")
}

func (a Agent) Business() {
	fmt.Println(a.start.getName()+"和"+a.company.getName()+"洽谈")
}