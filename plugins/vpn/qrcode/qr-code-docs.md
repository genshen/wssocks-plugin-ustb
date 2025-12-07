## Documentation of implementation of USTB vpn qr-code login

Starting from 2024/05, the ustb-vpn adds QR-code login.  
The authentication process can be described as following:  
1. In vpn login page's html code, (e.g., https://n.ustb.edu.cn/login), it contains an iframe of url:
   "https://sis.ustb.edu.cn/connect/qrpage?appid=****&return_url=https%3A%2F%2Fn.ustb.edu.cn%2Flogin%3Fustb_sis%3Dtrue&rand_token=***&embed_flag=1", 
   In the iframe, it contains a QR-code (qr code image url: https://sis.ustb.edu.cn/connect/qrimg?sid=${SID}, and QR content is:
   https://sis.ustb.edu.cn/auth?sid=${SID).
2. While waiting for WeChat QR scanning, it sends a requests of `https://sis.ustb.edu.cn/connect/state?sid=${SID}` (state request) and waits for its response.
3. If WeChat QR scanning is finished, the request in step 2. will response json data `{"state":200,"data":"39ff2a42e4474e70228b6337e159da8d"}`.
4. The `data` field in above json data is used as **auth_code**.
5. A callback redirect of url `https://n.ustb.edu.cn/login?ustb_sis=true&appid=***&auth_code=***&rand_token=***` will be generated and requested.
6. QR-code login finished.

In our go side implementation, we need:  
1) Parse _appid_, _rand\_token_, and _SID_ in the iframe url or its html content.
2) Generate QR code images: use _SID_ to generate the QR code by using package github.com/skip2/go-qrcode.
3) Send state request and waits for its response. After its response arrives, we can obtain the _auth_code_.
4) Send callback redirect request to finish QR-code login.
