// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: witnessupdateoperation.go

package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/denkhaus/bitshares/types"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *WitnessUpdateOperation) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *WitnessUpdateOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	if j.NewSigningKey != nil {
		buf.WriteString(`{ "new_signing_key":`)

		{

			obj, err = j.NewSigningKey.MarshalJSON()
			if err != nil {
				return err
			}
			buf.Write(obj)

		}
	} else {
		buf.WriteString(`{ "new_signing_key":null`)
	}
	buf.WriteString(`,"new_url":`)
	fflib.WriteJsonString(buf, string(j.NewURL))
	buf.WriteString(`,"witness":`)

	{

		obj, err = j.Witness.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"witness_account":`)

	{

		obj, err = j.WitnessAccount.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteByte(',')
	if j.Fee != nil {
		if true {
			/* Struct fall back. type=types.AssetAmount kind=struct */
			buf.WriteString(`"fee":`)
			err = buf.Encode(j.Fee)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtWitnessUpdateOperationbase = iota
	ffjtWitnessUpdateOperationnosuchkey

	ffjtWitnessUpdateOperationNewSigningKey

	ffjtWitnessUpdateOperationNewURL

	ffjtWitnessUpdateOperationWitness

	ffjtWitnessUpdateOperationWitnessAccount

	ffjtWitnessUpdateOperationFee
)

var ffjKeyWitnessUpdateOperationNewSigningKey = []byte("new_signing_key")

var ffjKeyWitnessUpdateOperationNewURL = []byte("new_url")

var ffjKeyWitnessUpdateOperationWitness = []byte("witness")

var ffjKeyWitnessUpdateOperationWitnessAccount = []byte("witness_account")

var ffjKeyWitnessUpdateOperationFee = []byte("fee")

// UnmarshalJSON umarshall json - template of ffjson
func (j *WitnessUpdateOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *WitnessUpdateOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtWitnessUpdateOperationbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtWitnessUpdateOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'f':

					if bytes.Equal(ffjKeyWitnessUpdateOperationFee, kn) {
						currentKey = ffjtWitnessUpdateOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'n':

					if bytes.Equal(ffjKeyWitnessUpdateOperationNewSigningKey, kn) {
						currentKey = ffjtWitnessUpdateOperationNewSigningKey
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyWitnessUpdateOperationNewURL, kn) {
						currentKey = ffjtWitnessUpdateOperationNewURL
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffjKeyWitnessUpdateOperationWitness, kn) {
						currentKey = ffjtWitnessUpdateOperationWitness
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyWitnessUpdateOperationWitnessAccount, kn) {
						currentKey = ffjtWitnessUpdateOperationWitnessAccount
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyWitnessUpdateOperationFee, kn) {
					currentKey = ffjtWitnessUpdateOperationFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyWitnessUpdateOperationWitnessAccount, kn) {
					currentKey = ffjtWitnessUpdateOperationWitnessAccount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyWitnessUpdateOperationWitness, kn) {
					currentKey = ffjtWitnessUpdateOperationWitness
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyWitnessUpdateOperationNewURL, kn) {
					currentKey = ffjtWitnessUpdateOperationNewURL
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyWitnessUpdateOperationNewSigningKey, kn) {
					currentKey = ffjtWitnessUpdateOperationNewSigningKey
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtWitnessUpdateOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtWitnessUpdateOperationNewSigningKey:
					goto handle_NewSigningKey

				case ffjtWitnessUpdateOperationNewURL:
					goto handle_NewURL

				case ffjtWitnessUpdateOperationWitness:
					goto handle_Witness

				case ffjtWitnessUpdateOperationWitnessAccount:
					goto handle_WitnessAccount

				case ffjtWitnessUpdateOperationFee:
					goto handle_Fee

				case ffjtWitnessUpdateOperationnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_NewSigningKey:

	/* handler: j.NewSigningKey type=types.PublicKey kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			j.NewSigningKey = nil

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			if j.NewSigningKey == nil {
				j.NewSigningKey = new(types.PublicKey)
			}

			err = j.NewSigningKey.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_NewURL:

	/* handler: j.NewURL type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.NewURL = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Witness:

	/* handler: j.Witness type=types.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Witness.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_WitnessAccount:

	/* handler: j.WitnessAccount type=types.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.WitnessAccount.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Fee:

	/* handler: j.Fee type=types.AssetAmount kind=struct quoted=false*/

	{
		/* Falling back. type=types.AssetAmount kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Fee)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
