package common

//CreateWechatPayCoreClient 创建为微信支付的client
//func CreateWechatPayCoreClient() (*core.Client, error) {
//
//	mchID, _ := env.MustGet("MCHID")                            // 商户号
//	mchCertificateSerialNumber, _ := env.MustGet("MCH_CER_NUM") // 商户证书序列号
//	mchAPIv3Key, _ := env.MustGet("MCH_API_V3_KEY")             // 商户APIv3密钥
//
//	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
//	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("../static/cert/wechat-native/apiclient_key.pem")
//	if err != nil {
//		log.Fatal("load merchant private key error")
//		return nil, err
//	}
//
//	ctx := context.Background()
//	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
//	opts := []core.ClientOption{
//		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
//	}
//	client, err := core.NewClient(ctx, opts...)
//	if err != nil {
//		log.Fatalf("new wechat pay client err:%s", err)
//		return nil, err
//	}
//
//	// 发送请求，以下载微信支付平台证书为例
//	// https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay5_1.shtml
//	//svc := certificates.CertificatesApiService{Client: client}
//	//resp, result, err := svc.DownloadCertificates(ctx)
//	//log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
//
//	return client, nil
//
//}
