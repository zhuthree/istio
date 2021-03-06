// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/adapter/servicecontrol/config/config.proto

/*
	Package config is a generated protocol buffer package.

	The `servicecontrol` adapter delivers logs and metrics to
	[Google Service Control](https://cloud.google.com/service-control).

	This adapter supports the [servicecontrolreport template](https://istio.io/docs/reference/config/policy-and-telemetry/templates/servicecontrolreport/),
	the [quota template](https://istio.io/docs/reference/config/policy-and-telemetry/templates/quota/),
	and the [apikey template](https://istio.io/docs/reference/config/policy-and-telemetry/templates/apikey/).

	It is generated from these files:
		mixer/adapter/servicecontrol/config/config.proto

	It has these top-level messages:
		RuntimeConfig
		Quota
		GcpServiceSetting
		Params
*/
package config

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import google_protobuf1 "github.com/gogo/protobuf/types"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Adapter runtime config paramters.
type RuntimeConfig struct {
	CheckCacheSize        int32                      `protobuf:"varint,1,opt,name=check_cache_size,json=checkCacheSize,proto3" json:"check_cache_size,omitempty"`
	CheckResultExpiration *google_protobuf1.Duration `protobuf:"bytes,2,opt,name=check_result_expiration,json=checkResultExpiration" json:"check_result_expiration,omitempty"`
}

func (m *RuntimeConfig) Reset()                    { *m = RuntimeConfig{} }
func (*RuntimeConfig) ProtoMessage()               {}
func (*RuntimeConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

type Quota struct {
	// Istio quota name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The corresponding Google quota metric name.
	GoogleQuotaMetricName string `protobuf:"bytes,2,opt,name=google_quota_metric_name,json=googleQuotaMetricName,proto3" json:"google_quota_metric_name,omitempty"`
	// Quota token expiration time period.
	Expiration *google_protobuf1.Duration `protobuf:"bytes,3,opt,name=expiration" json:"expiration,omitempty"`
}

func (m *Quota) Reset()                    { *m = Quota{} }
func (*Quota) ProtoMessage()               {}
func (*Quota) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{1} }

// Adapter setting for a managed GCP service.
type GcpServiceSetting struct {
	// Local service name on the mesh, which matches destination.service attribute.
	MeshServiceName string `protobuf:"bytes,1,opt,name=mesh_service_name,json=meshServiceName,proto3" json:"mesh_service_name,omitempty"`
	// Fully qualified GCP service name.
	GoogleServiceName string `protobuf:"bytes,2,opt,name=google_service_name,json=googleServiceName,proto3" json:"google_service_name,omitempty"`
	// Quota configs
	Quotas []*Quota `protobuf:"bytes,3,rep,name=quotas" json:"quotas,omitempty"`
}

func (m *GcpServiceSetting) Reset()                    { *m = GcpServiceSetting{} }
func (*GcpServiceSetting) ProtoMessage()               {}
func (*GcpServiceSetting) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{2} }

// Configuration format for the `servicecontrol` adapter.
//
// Sample adapter config:
//
// ```yaml
// apiVersion: "config.istio.io/v1alpha2"
// kind: servicecontrol
// metadata:
//   name: testhandler
//   namespace: istio-system
// spec:
//   runtime_config:
//     check_cache_size: 200
//     check_result_expiration: 60s
//   credential_path: "/path/to/token.json"
//   service_configs:
//     - mesh_service_name: "echo.local.svc"
//       google_service_name: "echo.endpoints.cloud.goog"
//       quotas:
//         - name: ratelimit.quota.istio-system
//           google_quota_metric_name: read-requests
//           expiration: 1m
// ```
type Params struct {
	RuntimeConfig *RuntimeConfig `protobuf:"bytes,1,opt,name=runtime_config,json=runtimeConfig" json:"runtime_config,omitempty"`
	// A path to JSON token file, usually mounted as Kubernetes secret on pod.
	CredentialPath string               `protobuf:"bytes,2,opt,name=credential_path,json=credentialPath,proto3" json:"credential_path,omitempty"`
	ServiceConfigs []*GcpServiceSetting `protobuf:"bytes,3,rep,name=service_configs,json=serviceConfigs" json:"service_configs,omitempty"`
}

func (m *Params) Reset()                    { *m = Params{} }
func (*Params) ProtoMessage()               {}
func (*Params) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{3} }

func init() {
	proto.RegisterType((*RuntimeConfig)(nil), "adapter.servicecontrol.config.RuntimeConfig")
	proto.RegisterType((*Quota)(nil), "adapter.servicecontrol.config.Quota")
	proto.RegisterType((*GcpServiceSetting)(nil), "adapter.servicecontrol.config.GcpServiceSetting")
	proto.RegisterType((*Params)(nil), "adapter.servicecontrol.config.Params")
}
func (m *RuntimeConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RuntimeConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.CheckCacheSize != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.CheckCacheSize))
	}
	if m.CheckResultExpiration != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.CheckResultExpiration.Size()))
		n1, err := m.CheckResultExpiration.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *Quota) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Quota) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.GoogleQuotaMetricName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.GoogleQuotaMetricName)))
		i += copy(dAtA[i:], m.GoogleQuotaMetricName)
	}
	if m.Expiration != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.Expiration.Size()))
		n2, err := m.Expiration.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *GcpServiceSetting) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GcpServiceSetting) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.MeshServiceName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.MeshServiceName)))
		i += copy(dAtA[i:], m.MeshServiceName)
	}
	if len(m.GoogleServiceName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.GoogleServiceName)))
		i += copy(dAtA[i:], m.GoogleServiceName)
	}
	if len(m.Quotas) > 0 {
		for _, msg := range m.Quotas {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintConfig(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RuntimeConfig != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.RuntimeConfig.Size()))
		n3, err := m.RuntimeConfig.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if len(m.CredentialPath) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.CredentialPath)))
		i += copy(dAtA[i:], m.CredentialPath)
	}
	if len(m.ServiceConfigs) > 0 {
		for _, msg := range m.ServiceConfigs {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintConfig(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RuntimeConfig) Size() (n int) {
	var l int
	_ = l
	if m.CheckCacheSize != 0 {
		n += 1 + sovConfig(uint64(m.CheckCacheSize))
	}
	if m.CheckResultExpiration != nil {
		l = m.CheckResultExpiration.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func (m *Quota) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.GoogleQuotaMetricName)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.Expiration != nil {
		l = m.Expiration.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func (m *GcpServiceSetting) Size() (n int) {
	var l int
	_ = l
	l = len(m.MeshServiceName)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.GoogleServiceName)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if len(m.Quotas) > 0 {
		for _, e := range m.Quotas {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	var l int
	_ = l
	if m.RuntimeConfig != nil {
		l = m.RuntimeConfig.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.CredentialPath)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if len(m.ServiceConfigs) > 0 {
		for _, e := range m.ServiceConfigs {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *RuntimeConfig) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RuntimeConfig{`,
		`CheckCacheSize:` + fmt.Sprintf("%v", this.CheckCacheSize) + `,`,
		`CheckResultExpiration:` + strings.Replace(fmt.Sprintf("%v", this.CheckResultExpiration), "Duration", "google_protobuf1.Duration", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Quota) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Quota{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`GoogleQuotaMetricName:` + fmt.Sprintf("%v", this.GoogleQuotaMetricName) + `,`,
		`Expiration:` + strings.Replace(fmt.Sprintf("%v", this.Expiration), "Duration", "google_protobuf1.Duration", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GcpServiceSetting) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GcpServiceSetting{`,
		`MeshServiceName:` + fmt.Sprintf("%v", this.MeshServiceName) + `,`,
		`GoogleServiceName:` + fmt.Sprintf("%v", this.GoogleServiceName) + `,`,
		`Quotas:` + strings.Replace(fmt.Sprintf("%v", this.Quotas), "Quota", "Quota", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Params) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params{`,
		`RuntimeConfig:` + strings.Replace(fmt.Sprintf("%v", this.RuntimeConfig), "RuntimeConfig", "RuntimeConfig", 1) + `,`,
		`CredentialPath:` + fmt.Sprintf("%v", this.CredentialPath) + `,`,
		`ServiceConfigs:` + strings.Replace(fmt.Sprintf("%v", this.ServiceConfigs), "GcpServiceSetting", "GcpServiceSetting", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringConfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *RuntimeConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RuntimeConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RuntimeConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CheckCacheSize", wireType)
			}
			m.CheckCacheSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CheckCacheSize |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CheckResultExpiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CheckResultExpiration == nil {
				m.CheckResultExpiration = &google_protobuf1.Duration{}
			}
			if err := m.CheckResultExpiration.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Quota) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Quota: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Quota: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoogleQuotaMetricName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GoogleQuotaMetricName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Expiration == nil {
				m.Expiration = &google_protobuf1.Duration{}
			}
			if err := m.Expiration.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GcpServiceSetting) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GcpServiceSetting: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GcpServiceSetting: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MeshServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MeshServiceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoogleServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GoogleServiceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quotas", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Quotas = append(m.Quotas, &Quota{})
			if err := m.Quotas[len(m.Quotas)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RuntimeConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RuntimeConfig == nil {
				m.RuntimeConfig = &RuntimeConfig{}
			}
			if err := m.RuntimeConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CredentialPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CredentialPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceConfigs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceConfigs = append(m.ServiceConfigs, &GcpServiceSetting{})
			if err := m.ServiceConfigs[len(m.ServiceConfigs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("mixer/adapter/servicecontrol/config/config.proto", fileDescriptorConfig)
}

var fileDescriptorConfig = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xe3, 0x95, 0x55, 0xc2, 0xd5, 0x5a, 0x6a, 0x98, 0x28, 0x93, 0xb0, 0xaa, 0x0a, 0x89,
	0x0a, 0x21, 0x67, 0x2a, 0x07, 0x40, 0xe2, 0xc4, 0x40, 0x9c, 0x40, 0x5b, 0x7a, 0x82, 0x8b, 0xe5,
	0xb9, 0x5e, 0x62, 0xd1, 0xc4, 0xc1, 0x71, 0xd0, 0xb4, 0x13, 0x1f, 0x00, 0x09, 0x3e, 0x06, 0x17,
	0xbe, 0xc7, 0x8e, 0x3b, 0x72, 0x24, 0xe1, 0xc2, 0x71, 0x1f, 0x01, 0xc5, 0x76, 0xa1, 0x11, 0xd2,
	0x7a, 0x8a, 0xf3, 0xde, 0xef, 0xfd, 0xdf, 0xff, 0xf9, 0x19, 0xee, 0xa7, 0xf2, 0x54, 0xe8, 0x90,
	0x2d, 0x58, 0x6e, 0x84, 0x0e, 0x0b, 0xa1, 0x3f, 0x4a, 0x2e, 0xb8, 0xca, 0x8c, 0x56, 0xcb, 0x90,
	0xab, 0xec, 0x44, 0xc6, 0xfe, 0x43, 0x72, 0xad, 0x8c, 0x42, 0x77, 0x3d, 0x4b, 0xda, 0x2c, 0x71,
	0xd0, 0xde, 0xad, 0x58, 0xc5, 0xca, 0x92, 0x61, 0x73, 0x72, 0x45, 0x7b, 0x38, 0x56, 0x2a, 0x5e,
	0x8a, 0xd0, 0xfe, 0x1d, 0x97, 0x27, 0xe1, 0xa2, 0xd4, 0xcc, 0x48, 0x95, 0xb9, 0xfc, 0xe4, 0x33,
	0x80, 0x3b, 0x51, 0x99, 0x19, 0x99, 0x8a, 0x03, 0xab, 0x83, 0xa6, 0xf0, 0x06, 0x4f, 0x04, 0x7f,
	0x4f, 0x39, 0xe3, 0x89, 0xa0, 0x85, 0x3c, 0x13, 0x23, 0x30, 0x06, 0xd3, 0xed, 0xa8, 0x6f, 0xe3,
	0x07, 0x4d, 0x78, 0x2e, 0xcf, 0x04, 0x3a, 0x82, 0xb7, 0x1d, 0xa9, 0x45, 0x51, 0x2e, 0x0d, 0x15,
	0xa7, 0xb9, 0x74, 0xe2, 0xa3, 0xad, 0x31, 0x98, 0xf6, 0x66, 0x77, 0x88, 0xeb, 0x4e, 0x56, 0xdd,
	0xc9, 0x0b, 0xdf, 0x3d, 0xda, 0xb5, 0x95, 0x91, 0x2d, 0x7c, 0xf9, 0xb7, 0x6e, 0xf2, 0x05, 0xc0,
	0xed, 0xa3, 0x52, 0x19, 0x86, 0x10, 0xbc, 0x96, 0xb1, 0xd4, 0xb5, 0xbe, 0x1e, 0xd9, 0x33, 0x7a,
	0x0c, 0x47, 0x4e, 0x90, 0x7e, 0x68, 0x18, 0x9a, 0x0a, 0xa3, 0x25, 0xa7, 0x96, 0xdb, 0xb2, 0xdc,
	0xae, 0xcb, 0x5b, 0x89, 0xd7, 0x36, 0xfb, 0xa6, 0x29, 0x7c, 0x0a, 0xe1, 0x9a, 0xb9, 0xce, 0x26,
	0x73, 0x6b, 0xf0, 0xe4, 0x3b, 0x80, 0xc3, 0x57, 0x3c, 0x9f, 0xbb, 0x3b, 0x9f, 0x0b, 0x63, 0x64,
	0x16, 0xa3, 0x07, 0x70, 0x98, 0x8a, 0x22, 0xa1, 0x7e, 0x15, 0x74, 0xcd, 0xea, 0xa0, 0x49, 0x78,
	0xdc, 0x36, 0x27, 0xf0, 0xa6, 0x77, 0xdd, 0xa2, 0x9d, 0xe1, 0xa1, 0x4b, 0xad, 0xf3, 0xcf, 0x60,
	0xd7, 0x8e, 0x57, 0x8c, 0x3a, 0xe3, 0xce, 0xb4, 0x37, 0xbb, 0x47, 0xae, 0x5c, 0x3c, 0xb1, 0xc3,
	0x46, 0xbe, 0x66, 0x52, 0x01, 0xd8, 0x3d, 0x64, 0x9a, 0xa5, 0x05, 0x9a, 0xc3, 0xbe, 0x76, 0xab,
	0xa5, 0x0e, 0xb5, 0x0e, 0x7b, 0xb3, 0x87, 0x1b, 0x04, 0x5b, 0xef, 0x21, 0xda, 0xd1, 0xad, 0xe7,
	0x71, 0x1f, 0x0e, 0xb8, 0x16, 0x0b, 0x91, 0x19, 0xc9, 0x96, 0x34, 0x67, 0x26, 0xf1, 0x93, 0xf4,
	0xff, 0x85, 0x0f, 0x99, 0x49, 0xd0, 0x5b, 0x38, 0x58, 0xcd, 0xeb, 0x74, 0x57, 0xf3, 0xec, 0x6f,
	0x68, 0xff, 0xdf, 0x6d, 0x47, 0x7d, 0x0f, 0x3a, 0x0b, 0xc5, 0xf3, 0x27, 0xe7, 0x15, 0x0e, 0x2e,
	0x2a, 0x1c, 0xfc, 0xa8, 0x70, 0x70, 0x59, 0xe1, 0xe0, 0x53, 0x8d, 0xc1, 0xb7, 0x1a, 0x07, 0xe7,
	0x35, 0x06, 0x17, 0x35, 0x06, 0x3f, 0x6b, 0x0c, 0x7e, 0xd7, 0x38, 0xb8, 0xac, 0x31, 0xf8, 0xfa,
	0x0b, 0x07, 0xef, 0xba, 0x4e, 0xfb, 0xb8, 0x6b, 0x97, 0xfd, 0xe8, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x89, 0x8a, 0x04, 0x79, 0x7e, 0x03, 0x00, 0x00,
}
