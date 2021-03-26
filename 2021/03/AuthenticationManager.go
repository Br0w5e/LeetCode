package main

//5694. 设计一个验证系统
type AuthenticationManager struct {
	Codes []Code
	TimeToLive int
}

type Code struct {
	TokenId string
	CurrentTime int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		TimeToLive: timeToLive,
	}
}


func (this *AuthenticationManager) Generate(tokenId string, currentTime int)  {
	code := Code{
		TokenId: tokenId,
		CurrentTime: currentTime,
	}
	this.Codes = append(this.Codes, code)
}


func (this *AuthenticationManager) Renew(tokenId string, currentTime int)  {
	for i := 0; i < len(this.Codes); i++ {
		if this.Codes[i].TokenId == tokenId && this.Codes[i].CurrentTime+this.TimeToLive > currentTime{
			this.Codes[i].CurrentTime = currentTime
		}
	}

}


func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	res := 0
	for i := 0; i < len(this.Codes); i++ {
		if this.Codes[i].CurrentTime+this.TimeToLive > currentTime {
			res++
		}
	}
	return res
}


/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */