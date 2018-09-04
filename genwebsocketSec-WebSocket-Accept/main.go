package main

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) >= 2 {
		fmt.Println(`Example: echo "dGhlIHNhbXBsZSBub25jZQ==" | genwebsocketSec-WebSocket-Accept `)
		fmt.Println(`will output: s3pPLMBiTxaQ9kYGzzhZRbK+xOo=  which is in http://www.rfc-editor.org/rfc/rfc6455.txt`)
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fromRequest := scanner.Text()
	all := fromRequest + "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
	rawbyte := Sha1(all)
	fmt.Println(base64.StdEncoding.EncodeToString(rawbyte))
}
func Sha1(in ...string) []byte {
	buf := new(bytes.Buffer)

	for idx := range in {
		buf.WriteString(in[idx])
	}
	tmp := sha1.Sum(buf.Bytes())

	return tmp[:]
}
