//
//  MinIO Inc [madmin-go]
//  Copyright (c) 2014-2025 MinIO.
//  All rights reserved. No warranty, explicit or implicit, provided.
//
package madmin

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *TierConfig) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Version":
			z.Version, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Version")
				return
			}
		case "Type":
			{
				var zb0002 int
				zb0002, err = dc.ReadInt()
				if err != nil {
					err = msgp.WrapError(err, "Type")
					return
				}
				z.Type = TierType(zb0002)
			}
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "S3":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "S3")
					return
				}
				z.S3 = nil
			} else {
				if z.S3 == nil {
					z.S3 = new(TierS3)
				}
				err = z.S3.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "S3")
					return
				}
			}
		case "Azure":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "Azure")
					return
				}
				z.Azure = nil
			} else {
				if z.Azure == nil {
					z.Azure = new(TierAzure)
				}
				err = z.Azure.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Azure")
					return
				}
			}
		case "GCS":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "GCS")
					return
				}
				z.GCS = nil
			} else {
				if z.GCS == nil {
					z.GCS = new(TierGCS)
				}
				err = z.GCS.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "GCS")
					return
				}
			}
		case "MinIO":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "MinIO")
					return
				}
				z.MinIO = nil
			} else {
				if z.MinIO == nil {
					z.MinIO = new(TierMinIO)
				}
				err = z.MinIO.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "MinIO")
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *TierConfig) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 7
	// write "Version"
	err = en.Append(0x87, 0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteString(z.Version)
	if err != nil {
		err = msgp.WrapError(err, "Version")
		return
	}
	// write "Type"
	err = en.Append(0xa4, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt(int(z.Type))
	if err != nil {
		err = msgp.WrapError(err, "Type")
		return
	}
	// write "Name"
	err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	// write "S3"
	err = en.Append(0xa2, 0x53, 0x33)
	if err != nil {
		return
	}
	if z.S3 == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.S3.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "S3")
			return
		}
	}
	// write "Azure"
	err = en.Append(0xa5, 0x41, 0x7a, 0x75, 0x72, 0x65)
	if err != nil {
		return
	}
	if z.Azure == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Azure.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Azure")
			return
		}
	}
	// write "GCS"
	err = en.Append(0xa3, 0x47, 0x43, 0x53)
	if err != nil {
		return
	}
	if z.GCS == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.GCS.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "GCS")
			return
		}
	}
	// write "MinIO"
	err = en.Append(0xa5, 0x4d, 0x69, 0x6e, 0x49, 0x4f)
	if err != nil {
		return
	}
	if z.MinIO == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.MinIO.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "MinIO")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TierConfig) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "Version"
	o = append(o, 0x87, 0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	o = msgp.AppendString(o, z.Version)
	// string "Type"
	o = append(o, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendInt(o, int(z.Type))
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "S3"
	o = append(o, 0xa2, 0x53, 0x33)
	if z.S3 == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.S3.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "S3")
			return
		}
	}
	// string "Azure"
	o = append(o, 0xa5, 0x41, 0x7a, 0x75, 0x72, 0x65)
	if z.Azure == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Azure.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Azure")
			return
		}
	}
	// string "GCS"
	o = append(o, 0xa3, 0x47, 0x43, 0x53)
	if z.GCS == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.GCS.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "GCS")
			return
		}
	}
	// string "MinIO"
	o = append(o, 0xa5, 0x4d, 0x69, 0x6e, 0x49, 0x4f)
	if z.MinIO == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.MinIO.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "MinIO")
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TierConfig) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Version":
			z.Version, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Version")
				return
			}
		case "Type":
			{
				var zb0002 int
				zb0002, bts, err = msgp.ReadIntBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Type")
					return
				}
				z.Type = TierType(zb0002)
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "S3":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.S3 = nil
			} else {
				if z.S3 == nil {
					z.S3 = new(TierS3)
				}
				bts, err = z.S3.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "S3")
					return
				}
			}
		case "Azure":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Azure = nil
			} else {
				if z.Azure == nil {
					z.Azure = new(TierAzure)
				}
				bts, err = z.Azure.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Azure")
					return
				}
			}
		case "GCS":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.GCS = nil
			} else {
				if z.GCS == nil {
					z.GCS = new(TierGCS)
				}
				bts, err = z.GCS.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "GCS")
					return
				}
			}
		case "MinIO":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.MinIO = nil
			} else {
				if z.MinIO == nil {
					z.MinIO = new(TierMinIO)
				}
				bts, err = z.MinIO.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "MinIO")
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TierConfig) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.Version) + 5 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.Name) + 3
	if z.S3 == nil {
		s += msgp.NilSize
	} else {
		s += z.S3.Msgsize()
	}
	s += 6
	if z.Azure == nil {
		s += msgp.NilSize
	} else {
		s += z.Azure.Msgsize()
	}
	s += 4
	if z.GCS == nil {
		s += msgp.NilSize
	} else {
		s += z.GCS.Msgsize()
	}
	s += 6
	if z.MinIO == nil {
		s += msgp.NilSize
	} else {
		s += z.MinIO.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TierType) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 int
		zb0001, err = dc.ReadInt()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = TierType(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z TierType) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt(int(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z TierType) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt(o, int(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TierType) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 int
		zb0001, bts, err = msgp.ReadIntBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = TierType(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z TierType) Msgsize() (s int) {
	s = msgp.IntSize
	return
}
