// AUTOGENERATED CODE. DO NOT EDIT.

package gendoc

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
)

var embeddedResources = map[string]string{
	"docbook.tmpl": "H4sIAAAAAAAA/+xZ30/bOhR+719h5fEiEq64SFeTW6TBqmkChIDt3U1OW2uOndlOAUX93ycnIb8Tl7UUtvGCZJ/P58TH5/uOa/DpQ8jQCqSigo+df90jBwH3RUD5Yux8vZse/u+cTkaYSE19BpMRQlhTzWByLYUWvmDoXPhxCFwTTQXHXmYdIZQkkvAFIHdKGaj12ixV4BuUMReOksS9IiGs15W1ZnVEJEHuOShf0sisSl1U/F6CUmSRuy6dIxqMnSRxpzFjmWMnc1mNeCH4oiPqUFxjo3PkfiZqSoEFqpjHmswYoLkkIYwdwlgRsAiJfUaU4iRsRS8NKHPb+CDjYiFFHCFfMDV2/qs4RwibyQh8Y7yngV6OnX8cb2vEkXtiBx03IXoJJKjOIISluK/PIISBa/k4SXeLvWzQDbl7jGAYcUFmwIYhlZPsBGKv8Y3Ya20E65kIGusq9V2rBuvGKwU/9N2YUf4dmT/Ay4o2KTEVnVdRNsSegU2G/ZkVJlu2uJ0MyEr/HOYkZvobYbGJanCTfO4DSpKm3UsBSQI86Anayr1Jawqvn0c9+9jLGFHQ2ksJWFK46qEg7acHDdzI3O6JewVKQ4DKCBYOn+yDw2+D5UVOtmX6R6IsiKs4nIF8ZTHoqDJrjl5JENr+zgTXhHLKFw3PpeH5opMdy1+lOtirXXSqprJQeBzu7+qyK6VLk2yTt+N9yNuWumT29hvISZbvnUvJdrTcF6l6aJQP+tp6/UdGUd7mV8ghgxWw/j6NKZ8LGRI2yJb3Tv7eyd87+R/VyWu830R88hq5Bbmi/q89Qbx8D+/o35egl+JtvDG8uGBle0X2Tn8DP2JQGtml6wZUJLiCDaAvrk/5Ue5BnPL8NJQkn91aqp5y2nKfTT9bpt7ehaVhqEhJ54PorU8YkdlVOy20Ol3tl5T+K4qNhh32Extgt/Y2B1qnlJ+0G0mhRT8Tn+4SQpsE9gPODg6sTr6QFbGCrh/1UvA+WKPcWmxvc73sMmlB1Jnel5an1pM+z1foVBkP7cFohEmYFXUWRRt5M5nbCJhlrxfaYmuTqw2m1nna0eErlBxhr/j/xs8AAAD//2FyhLQRGQAA",
	"html.tmpl": "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",
	"markdown.tmpl": "H4sIAAAAAAAA/+RWz2/aMBS+5694IzusqgL3CjisXTVNbVW1aJdqWg08IJKxs9hBq2L/75N/JHaADKRKuyyX+Nl5L9/7vs9OUngsueQLTuGGL6otMklkzlkyJsDIFicDyYvBaJokaQozMqcIfAXXnElkUiR1XRK2Rhje5hSF1kldf1zlFH+aXLiawPCBbFHrDF7q2o9/fErb8UUC0Na4RyHI2pQBcAl3nK3jpNuK0jgR2VLruMQXVm3flf9bIhM5Z00R01dGcYcUwpotFvrUOsN2rafwM5a7fBH3dhpXc8/g5XlBKCnhO6EVwuytQINB2MlsZyYzaSYvkrMVCQK3WIzM4wIIzddsMijz9UYOpmMCmxJXk0FqjDCd8WI8ItPxqHCWaJOTuh7eoFiUeWH8o3UEJQjbeWvo2hvMlguaHS2Zr2D4lYjbHOnSFFRgh6AsLaDgjsyRgoIoE1SiIDMXuDt0Q3+Bitkz9SGzKqrQJqjWWOZ9sYAuvnBPWxT28bpmfF5CtxPXxw2uSEWlVVVr8OEV2L7jJe8EC7DrjcBIbN1EBbcGZj4TYW4P1XaOZR9Dhyy1JPWzFd59lLEOYS42JwjJWc7W+ysO3r+hzi2NP2QZIFvC1vsUsmwa2dcfKu/1rgKzdEqAM9i2nfUyfYq/iIlDBtC02rR/zFgRBz0HYMvH0aPz/3Xm64E1X3u9GQTp8L/ny/BNOcOaf7HlPcoNXzbufMJfFQrZSPOEouBMYBP3SrOvwn64H6v4+2AA9J+2HtL+oeunu2eve95hPivh1CY5kER42hs1/F5KUzj8Sht6h4X5w2r4e+ASBSi4vrxspr6RHWnGj29yw5mPjtMb8QiHW+PIHnCwAr2xWezfnyNjAKMpdKe8FgZyuz2KIl4z2OPY4W9mWnb+BAAA///ASSPYZwoAAA==",
	"scalars.json": "H4sIAAAAAAAA/9yXzW4aMRDH7zzFiFMqBZDSlEa9JZGQOOQEOUWp5GVnvW6NTewxzaqq1HfoG/ZJqt0F1gYvpChVk9zQfHg9v/nPWNx1AL53AAC6C6NJT4sFdj9BN9Uukdg9rV1KE9rSvDbMFot45MzG7XxzciY1o+H52vGFLVk8ZZEvgpyNvaBcq6jLuKRYO0aVowPw47SlxiB1X4lBYFNhYN4q8P1ZrMCwjn9dn1DeNTb13Vq0sGRGsERiT6LilAOqmU6F4n0YK8wyMROoCDJtNh5QyBmJJYJy8wSNhd8/f4HIoNDOQCZQpiAsSPEVZQGkIWdl7DppyaRDewrOItjqYiCUJWRpPwI8uHkDXKgI7iDWh+1Fe6iFIuRo4rC9FB/1leDKzUEbGInH8tcJs2DwwQmD6btDPWi0/sJ6MDw/0IPm5k0PpFY82oT4SPvhu10YWDLCC9huxsBP3+3IXvLuGPlHOLgWMbqoGt0zy/EAgaM06Y4RZQuZmERcXCPu/4jkKET2qN05EVxhCkJRPWt9mOZoEebaIGxGWhZ1Cu6OM+VMgUHuJDNQXcG+7fVoj9qPz855eN7G+VWvwEw8YhqR8aX8xgoLWflqJAWh7cNNgK4G5FbvdLZ6OYAZBJ0RKuAGGaGp484+n128vM35zEqtWEakumKJgud0CGb54B6G+WH4epft/mF/uh7f9tL7Cy29wa2UaC2f8q/Lj2vK9K08eqZfZWlHpiKFbnuCCrecfoFT4/BaMmsHIyZt/XN/twOSTa+hdsDcWQJWd36mFTGh4HY66l2sXq+01NjHXiIILifX4zEQPlJMF+GHGmKhnbdczKc2CZuft3wiZGbJDJwS5ZVj3Ooz4aQqbf98VMrfAXbDig0fpgpgJhFkmCnA4oNDNSvXafvUtNG5KggnLYTu7svjYoR2s55OaR+dqsO9i6vxtEbUue/8CQAA//8z+wC/ohEAAA==",
}

func fetchResource(name string) ([]byte, error) {
	raw, ok := embeddedResources[name]
	if !ok {
		return nil, fmt.Errorf("Could not find resource for '%s'", name)
	}

	compressed, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	buf := bytes.NewBuffer(compressed)
	
	r, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(&out, r); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
