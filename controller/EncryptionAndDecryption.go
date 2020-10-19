package controller

import (
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/yanlong-li/hi-go-server/packet_model/encrypt"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/packet"
	"github.com/yanlong-li/hi-go-socket/stream"
)

//加密和解密
var encryptedWhiteList = make(map[uint32]bool, 0)

func init() {

	encryptedWhiteList[7001] = true
	encryptedWhiteList[7002] = true

	//route.Register(packet_model.BeforeSending{}, Encryption)
	//route.Register(encrypt.BytesData{}, Decryption)
}

// 加密动作
func Encryption(ps stream.Interface, conn connect.Connector) []byte {

	var data []byte
	if _, ok := encryptedWhiteList[ps.GetOpCode()]; !ok {
		// 加密之前要先判断是否需要加密，比如 协商加密协议的数据 发送基础配置的数据 不需要加密
		c, e := helper.EncryptMessageArmored("-----BEGIN PGP PUBLIC KEY BLOCK-----\nComment: 用户编号:\ttesting <test@test.test>\nComment: 连接时间:\t2020/4/10 15:24\nComment: 过期时间:\t2022/4/10 12:00\nComment: 文件系统:\t512 位 ECDSA\nComment: 用途:\t签名, 仅加密, 正在认证的用户编号, SSH 认证\nComment: 指纹:\t8FF1A48D9B97C67F9B8C4F633708AAA986797D5D\n\n\nmJMEXpAfNBMJKyQDAwIIAQENBAMETjfSIfPWeqGQZZrKDce9RDkI0FlqmgvusIST\n/QigTSqi+Zy30hZBAFV/4Xrd1iSkUtj4IRxr+raJhN3cJ0ga3FFyZZLAladulVWs\nm16Kzurj80s2AkRYu9Dh+ghACPH/d430EnW+CUR/DoLpFyNgBYBkFi0qayGX8uii\naGSZg4q0GHRlc3RpbmcgPHRlc3RAdGVzdC50ZXN0PojWBBMTCgA+FiEEj/GkjZuX\nxn+bjE9jNwiqqYZ5fV0FAl6QHzQCGyMFCQPCNwwFCwkIBwIGFQoJCAsCBBYCAwEC\nHgECF4AACgkQNwiqqYZ5fV39WgIAjtDn+fwR2mPojJUXZzQiohS++z43Ll6/kPkN\nd3l8MD87UcohDYC2bmSFsMKzIcLaU9i0cJUVWjD1hdlJOQEBzgH+MumW3QcN7g2y\nHxkxWLd0Wr2GQIx56kf/2x+M/ANwNF86ggwmmo8DAI+5QtYxE1DBCllaXNvbh8Ag\neTX83XnApbiXBF6QHzQSCSskAwMCCAEBDQQDBG02qIpJ/PL3KY7r1hQMjAWqXd1w\njI0hpxbSKu+EOu34Ybx6nyTQz4pACT9lMeVvj0nzZd6K6zHRUTB+0JlA5k5BavFz\nd8J25BKBFW20A7IhKMv3NnRcfQoL5gqvu7rMUlfxRyfpRpStkhTxtwrPhScv6Uu4\nBOkSjFHp2e/+OjKdAwEKCYi+BBgTCgAmFiEEj/GkjZuXxn+bjE9jNwiqqYZ5fV0F\nAl6QHzQCGwwFCQPCNwwACgkQNwiqqYZ5fV16bwH/dBuAFp8+TVPm8lX77KunTrWN\nIHzIfaeB5hUPIeU9N3m4EQXXLz5HKl1Tww/G4He/paWMu3QpN9+AABQjQp1ARwIA\nnEdgVA7npVfaZScOMV1mguzTeu6VV+3xgHppB1C/ssp/4+M6QvkD8DBw9T7xEsu7\niElz4OMBeZf0+Ugj+iLDWg==\n=dxC2\n-----END PGP PUBLIC KEY BLOCK-----\n", string(ps.ToData()))
		if e != nil {
			logger.Warning("加密失败", 0, e)
		}
		data = []byte(c)

		//简单做了一层封装,告知对方这是一个指定类型的数据包 ， 你可以定义其他类型比如增加字段 时间戳 或者不加密而是签名 encrypt.BytesData{Data: data, Sign: sign}
		model := encrypt.BytesData{Data: data}
		ps.Marshal(model)
		// 返回加密后的数据
		return ps.ToData()
	} else {
		// 原样返回
		return ps.ToData()
	}
}
func Decryption(encryptData encrypt.BytesData, conn connect.Connector) {

	logger.Debug("收到一个加密数据", 0)

	if len(encryptData.Data) < packet.OpCodeLen+packet.BufLenLen {
		logger.Warning("解密后的数据无效", 0, encryptData.Data)
	}
	//自行处理粘包和拆包，单个包处理后丢给 conn.HandleData()

	//bufLen := buf[:packet.BufLenLen]
	bufData := encryptData.Data[packet.BufLenLen:]
	conn.HandleData(bufData)
}
