package common

/**CreateClient
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
//func CreateClient() (_result *ocr20191230.Client, _err error) {
//	accessKeyId, _ := env.MustGet("OCR_ACCESS_KEY_ID")
//	accessKeySecret, _ := env.MustGet("OCR_ACCESS_SECRET")
//	ocrEndPoint, _ := env.MustGet("OCR_END_POINT")
//
//	config := &openapi.Config{
//		// 您的AccessKey ID
//		AccessKeyId: tea.String(accessKeyId),
//		// 您的AccessKey Secret
//		AccessKeySecret: tea.String(accessKeySecret),
//	}
//	// 访问的域名
//	config.Endpoint = tea.String(ocrEndPoint)
//	_result = &ocr20191230.Client{}
//	_result, _err = ocr20191230.NewClient(config)
//	return _result, _err
//}
