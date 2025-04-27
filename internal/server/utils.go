package server

import (
	"bufio"
	"fmt"
)

func parseRESP(reader *bufio.Reader) ([]string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	if len(line) < 1 || line[0] != '*' {
		return nil, fmt.Errorf("expected array, got %q", line)
	}

	var count int
	if _, err := fmt.Sscanf(line, "*%d\r\n", &count); err != nil {
		return nil, fmt.Errorf("malformed array header: %w", err)
	}

	cmd := make([]string, 0, count)
	for i := 0; i < count; i++ {
		typ, err := reader.ReadByte()
		if err != nil {
			return nil, err
		}
		if typ != '$' {
			return nil, fmt.Errorf("expected bulk string, got %q", typ)
		}
		lenLine, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		var strLen int
		if _, err := fmt.Sscanf(lenLine, "%d\r\n", &strLen); err != nil {
			return nil, fmt.Errorf("malformed bulk length: %w", err)
		}

		buf := make([]byte, strLen+2)
		if _, err := reader.Read(buf); err != nil {
			return nil, err
		}
		cmd = append(cmd, string(buf[:strLen]))
	}
	return cmd, nil
}

func generateResponse(resp Response) string {
	var data string
	if resp.IsNull {
		data = "$-1\r\n"
		return data
	}
	if resp.IsArrayResponse {
		data = generateArrayResponse(resp.Array)
	} else {
		data = fmt.Sprintf("$%d\r\n%s\r\n", len(resp.Body), resp.Body)
	}
	return data
}

func generateArrayResponse(resp []string) string {
	data := fmt.Sprintf("*%d\r\n", len(resp))
	for _, bulkString := range resp {
		size := len(bulkString)
		suffix := fmt.Sprintf("$%d\r\n%s\r\n", size, bulkString)
		data += suffix
	}
	return data
}
