package main

type VkToken struct {
	Id      uint32 `json:"user_id"`
	Profile int    `json:"profile_type"`
	Token   string `json:"access_token"`
}
type Tokens []VkToken

type VkProfileResponse struct {
	Response struct {
		Profiles []struct {
			Id         uint32 `json:"id"`
			Eduprofile struct {
				Eduroles struct {
					Mainroles     string `json:"main_roles"`
					Organizations []struct {
						Name string `json:"organization_name"`
					} `json:"organizations"`
				} `json:"edu_roles"`
			} `json:"educational_profile"`
			Name    string `json:"first_name"`
			Surname string `json:"last_name"`
		} `json:"profiles"`
	} `json:"response"`
}

func findSferumToken(tokens Tokens) VkToken {
	for _, t := range tokens {
		if t.Profile == 2 {
			return t
		}
	}
	return VkToken{}
}
