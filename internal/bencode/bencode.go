package bencode

import (
	"fmt"
	"sort"
	"strings"
)

/*
Bencode:
- Strings are length-prefixed base ten followed by a colon and the string. For example 4:spam corresponds to 'spam'.
- Integers are represented by an 'i' followed by the number in base 10 followed by an 'e'. For example i3e corresponds to 3 and i-3e corresponds to -3. Integers have no size limiation. i-0e is invalid. All encodings with a leading zero, such as i03e, are invalid, other than i0e which of course corresponds to 0.
- Lists are encoded as an 'l' followed by their elements (also bencoded) followed by an 'e'. For example l4:spam4:eggse correspons to ['spam','eggs'].
- Dictionaries are encoded as a 'd' followed by a list of alternating keys and their corresponding values followed by an 'e'. For example, d3:cow3:moo4:spam4:eggse corresponds to {'cow':'moo', 'spam':'eggs'} and d4:spam11:a1:bee corresponds to {'spam':['a':'b']}. Keys must be strings and appear in sorted order.
*/

func EncodeBencode(data any) (string, error) {
	switch v:=data.(type) {

	case string:
		return fmt.Sprintf("%d:%s", len(v), v), nil

	case int:
		return fmt.Sprintf("i%de", v), nil

	case []any:
		var b strings.Builder
		b.WriteByte('l')
		
		for _, item := range v {
			encodedItem, err := EncodeBencode(item)
			if err != nil {
				return "", err
			}
			b.WriteString(encodedItem)
		}

		b.WriteByte('e')	
		return b.String(), nil

	case map[string]any:
		var b strings.Builder
		b.WriteByte('d')

		keys := make([]string, 0, len(v))
		for key := range v {
			keys = append(keys, key)
		}
		sort.Strings(keys) // Bencode requires dict to be sorted in lexicographical order.

		for _, key := range keys {
			b.WriteString(fmt.Sprintf("%d:%s", len(key), key))
			encodedValue, err := EncodeBencode(v[key])
			if err != nil {
				return "", err
			}
			b.WriteString(encodedValue)
		}

		b.WriteByte('e')
		return b.String(), nil
	default:
		return "", fmt.Errorf("unsupported type for bencode encoding: %T", data)
	}
}

/*
func decodeBencode(bencodedString string) (interface{}, int, error) {

}
*/
