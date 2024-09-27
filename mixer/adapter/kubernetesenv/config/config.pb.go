// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/adapter/kubernetesenv/config/config.proto

/*
	Package config is a generated protocol buffer package.

	The `kubernetesenv` adapter extracts information from a Kubernetes environment
	and produces attribtes that can be used in downstream adapters.

	This adapter supports the [kubernetesenv template](https://istio.io/docs/reference/config/policy-and-telemetry/templates/kubernetes/).

	It is generated from these files:
		mixer/adapter/kubernetesenv/config/config.proto

	It has these top-level messages:
		Params
*/
package config

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"

import time "time"

import types "github.com/gogo/protobuf/types"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Configuration parameters for the kubernetes adapter. These params
// control the manner in which the kubernetes adapter discovers and
// generates values related to pod information.
//
// The adapter works by looking up pod information by UIDs (of the
// form: "kubernetes://pod.namespace"). It expects that the UIDs will be
// supplied in an input map for three distinct traffic classes (source,
// destination, and origin).
//
// For all valid UIDs supplied, this adapter generates output
// values containing information about the related pods.
type Params struct {
	// File path to discover kubeconfig. For in-cluster configuration,
	// this should be left unset. For local configuration, this should
	// be set to the path of a kubeconfig file that can be used to
	// reach a kubernetes API server.
	//
	// NOTE: The kubernetes adapter will use the value of the env var
	// KUBECONFIG in the case where it is set (overriding any value configured
	// through this proto).
	//
	// Default: "" (unset)
	KubeconfigPath string `protobuf:"bytes,1,opt,name=kubeconfig_path,json=kubeconfigPath,proto3" json:"kubeconfig_path,omitempty"`
	// Controls the resync period of the kubernetes cluster info cache.
	// The cache will watch for events and every so often completely resync.
	// This controls how frequently the complete resync occurs.
	//
	// Default: 5 minutes
	CacheRefreshDuration time.Duration `protobuf:"bytes,2,opt,name=cache_refresh_duration,json=cacheRefreshDuration,stdduration" json:"cache_refresh_duration"`
	// Configures the cluster domain name to use for service name normalization.
	//
	// Default: svc.cluster.local
	ClusterDomainName string `protobuf:"bytes,3,opt,name=cluster_domain_name,json=clusterDomainName,proto3" json:"cluster_domain_name,omitempty"`
	// In order to extract the service associated with a source, destination, or
	// origin, this adapter relies on pod labels. In particular, it looks for
	// the value of a specific label, as specified by this parameter.
	//
	// Default: app
	PodLabelForService string `protobuf:"bytes,4,opt,name=pod_label_for_service,json=podLabelForService,proto3" json:"pod_label_for_service,omitempty"`
	// In order to extract the service associated with a source, destination, or
	// origin, this adapter relies on pod labels. In particular, it looks for
	// the value of a specific label for istio component services, as specified
	// by this parameter.
	//
	// Default: istio
	PodLabelForIstioComponentService string `protobuf:"bytes,5,opt,name=pod_label_for_istio_component_service,json=podLabelForIstioComponentService,proto3" json:"pod_label_for_istio_component_service,omitempty"`
	//
	// Default: false
	LookupIngressSourceAndOriginValues bool `protobuf:"varint,6,opt,name=lookup_ingress_source_and_origin_values,json=lookupIngressSourceAndOriginValues,proto3" json:"lookup_ingress_source_and_origin_values,omitempty"`
	// Istio ingress service string. This is used to identify the
	// ingress service in requests.
	//
	// Default: "ingress.istio-system.svc.cluster.local"
	FullyQualifiedIstioIngressServiceName string `protobuf:"bytes,7,opt,name=fully_qualified_istio_ingress_service_name,json=fullyQualifiedIstioIngressServiceName,proto3" json:"fully_qualified_istio_ingress_service_name,omitempty"`
}

func (m *Params) Reset()                    { *m = Params{} }
func (*Params) ProtoMessage()               {}
func (*Params) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func init() {
	proto.RegisterType((*Params)(nil), "adapter.kubernetesenv.config.Params")
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
	if len(m.KubeconfigPath) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.KubeconfigPath)))
		i += copy(dAtA[i:], m.KubeconfigPath)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintConfig(dAtA, i, uint64(types.SizeOfStdDuration(m.CacheRefreshDuration)))
	n1, err := types.StdDurationMarshalTo(m.CacheRefreshDuration, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.ClusterDomainName) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ClusterDomainName)))
		i += copy(dAtA[i:], m.ClusterDomainName)
	}
	if len(m.PodLabelForService) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.PodLabelForService)))
		i += copy(dAtA[i:], m.PodLabelForService)
	}
	if len(m.PodLabelForIstioComponentService) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.PodLabelForIstioComponentService)))
		i += copy(dAtA[i:], m.PodLabelForIstioComponentService)
	}
	if m.LookupIngressSourceAndOriginValues {
		dAtA[i] = 0x30
		i++
		if m.LookupIngressSourceAndOriginValues {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.FullyQualifiedIstioIngressServiceName) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.FullyQualifiedIstioIngressServiceName)))
		i += copy(dAtA[i:], m.FullyQualifiedIstioIngressServiceName)
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
func (m *Params) Size() (n int) {
	var l int
	_ = l
	l = len(m.KubeconfigPath)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = types.SizeOfStdDuration(m.CacheRefreshDuration)
	n += 1 + l + sovConfig(uint64(l))
	l = len(m.ClusterDomainName)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.PodLabelForService)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.PodLabelForIstioComponentService)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.LookupIngressSourceAndOriginValues {
		n += 2
	}
	l = len(m.FullyQualifiedIstioIngressServiceName)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
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
func (this *Params) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params{`,
		`KubeconfigPath:` + fmt.Sprintf("%v", this.KubeconfigPath) + `,`,
		`CacheRefreshDuration:` + strings.Replace(strings.Replace(this.CacheRefreshDuration.String(), "Duration", "google_protobuf1.Duration", 1), `&`, ``, 1) + `,`,
		`ClusterDomainName:` + fmt.Sprintf("%v", this.ClusterDomainName) + `,`,
		`PodLabelForService:` + fmt.Sprintf("%v", this.PodLabelForService) + `,`,
		`PodLabelForIstioComponentService:` + fmt.Sprintf("%v", this.PodLabelForIstioComponentService) + `,`,
		`LookupIngressSourceAndOriginValues:` + fmt.Sprintf("%v", this.LookupIngressSourceAndOriginValues) + `,`,
		`FullyQualifiedIstioIngressServiceName:` + fmt.Sprintf("%v", this.FullyQualifiedIstioIngressServiceName) + `,`,
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
				return fmt.Errorf("proto: wrong wireType = %d for field KubeconfigPath", wireType)
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
			m.KubeconfigPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CacheRefreshDuration", wireType)
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
			if err := types.StdDurationUnmarshal(&m.CacheRefreshDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterDomainName", wireType)
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
			m.ClusterDomainName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PodLabelForService", wireType)
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
			m.PodLabelForService = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PodLabelForIstioComponentService", wireType)
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
			m.PodLabelForIstioComponentService = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LookupIngressSourceAndOriginValues", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.LookupIngressSourceAndOriginValues = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FullyQualifiedIstioIngressServiceName", wireType)
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
			m.FullyQualifiedIstioIngressServiceName = string(dAtA[iNdEx:postIndex])
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
	proto.RegisterFile("mixer/adapter/kubernetesenv/config/config.proto", fileDescriptorConfig)
}

var fileDescriptorConfig = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x4d, 0x6b, 0xd4, 0x40,
	0x18, 0xc7, 0x33, 0x56, 0xd7, 0x1a, 0x41, 0x31, 0x56, 0x59, 0x8b, 0x4c, 0x97, 0x42, 0xe9, 0xe2,
	0x21, 0x41, 0xbd, 0x78, 0xb5, 0x16, 0xa1, 0x20, 0xb6, 0xee, 0x82, 0x50, 0x2f, 0xc3, 0x6c, 0xf2,
	0x24, 0x3b, 0x34, 0x99, 0x27, 0xce, 0xcb, 0xa2, 0x37, 0x3f, 0x82, 0x47, 0x3f, 0x82, 0x1f, 0x65,
	0x8f, 0x3d, 0x7a, 0x52, 0x37, 0x5e, 0x3c, 0xee, 0x47, 0x90, 0xcc, 0x24, 0xb5, 0x3d, 0xe5, 0xe5,
	0xff, 0x7b, 0x7e, 0xff, 0x87, 0x99, 0x30, 0xa9, 0xc4, 0x27, 0x50, 0x09, 0xcf, 0x78, 0x6d, 0x40,
	0x25, 0x67, 0x76, 0x06, 0x4a, 0x82, 0x01, 0x0d, 0x72, 0x91, 0xa4, 0x28, 0x73, 0x51, 0x74, 0x8f,
	0xb8, 0x56, 0x68, 0x30, 0x7a, 0xdc, 0xa1, 0xf1, 0x15, 0x34, 0xf6, 0xcc, 0xf6, 0x56, 0x81, 0x05,
	0x3a, 0x30, 0x69, 0xdf, 0xfc, 0xcc, 0x36, 0x2d, 0x10, 0x8b, 0x12, 0x12, 0xf7, 0x35, 0xb3, 0x79,
	0x92, 0x59, 0xc5, 0x8d, 0x40, 0xe9, 0xf3, 0xdd, 0xf5, 0x46, 0x38, 0x38, 0xe1, 0x8a, 0x57, 0x3a,
	0xda, 0x0f, 0xef, 0xb6, 0x62, 0xaf, 0x63, 0x35, 0x37, 0xf3, 0x21, 0x19, 0x91, 0xf1, 0xad, 0xc9,
	0x9d, 0xff, 0xbf, 0x4f, 0xb8, 0x99, 0x47, 0xa7, 0xe1, 0xc3, 0x94, 0xa7, 0x73, 0x60, 0x0a, 0x72,
	0x05, 0x7a, 0xce, 0x7a, 0xe7, 0xf0, 0xda, 0x88, 0x8c, 0x6f, 0x3f, 0x7b, 0x14, 0xfb, 0xd2, 0xb8,
	0x2f, 0x8d, 0x0f, 0x3b, 0xe0, 0x60, 0x73, 0xf9, 0x73, 0x27, 0xf8, 0xf6, 0x6b, 0x87, 0x4c, 0xb6,
	0x9c, 0x62, 0xe2, 0x0d, 0x7d, 0x1e, 0xc5, 0xe1, 0xfd, 0xb4, 0xb4, 0xda, 0x80, 0x62, 0x19, 0x56,
	0x5c, 0x48, 0x26, 0x79, 0x05, 0xc3, 0x0d, 0xb7, 0xc7, 0xbd, 0x2e, 0x3a, 0x74, 0xc9, 0x5b, 0x5e,
	0x41, 0xf4, 0x34, 0x7c, 0x50, 0x63, 0xc6, 0x4a, 0x3e, 0x83, 0x92, 0xe5, 0xa8, 0x98, 0x06, 0xb5,
	0x10, 0x29, 0x0c, 0xaf, 0xbb, 0x89, 0xa8, 0xc6, 0xec, 0x4d, 0x9b, 0xbd, 0x46, 0x35, 0xf5, 0x49,
	0x74, 0x1c, 0xee, 0x5d, 0x1d, 0x11, 0xda, 0x08, 0x64, 0x29, 0x56, 0x35, 0x4a, 0x90, 0xe6, 0x42,
	0x71, 0xc3, 0x29, 0x46, 0x97, 0x14, 0x47, 0x2d, 0xf9, 0xaa, 0x07, 0x7b, 0xe1, 0x34, 0xdc, 0x2f,
	0x11, 0xcf, 0x6c, 0xcd, 0x84, 0x2c, 0x14, 0x68, 0xcd, 0x34, 0x5a, 0x95, 0x02, 0xe3, 0x32, 0x63,
	0xa8, 0x44, 0x21, 0x24, 0x5b, 0xf0, 0xd2, 0x82, 0x1e, 0x0e, 0x46, 0x64, 0xbc, 0x39, 0xd9, 0xf5,
	0xf8, 0x91, 0xa7, 0xa7, 0x0e, 0x7e, 0x29, 0xb3, 0x63, 0x87, 0xbe, 0x77, 0x64, 0x74, 0x1a, 0x3e,
	0xc9, 0x6d, 0x59, 0x7e, 0x66, 0x1f, 0x2d, 0x2f, 0x45, 0x2e, 0x20, 0xeb, 0xf6, 0xbc, 0xe8, 0xf0,
	0xed, 0xfe, 0x7c, 0x6e, 0xba, 0x55, 0xf7, 0xdc, 0xc4, 0xbb, 0x7e, 0xc0, 0x6d, 0xdb, 0x97, 0x78,
	0xba, 0x3d, 0xb3, 0x83, 0x17, 0xcb, 0x15, 0x0d, 0xce, 0x57, 0x34, 0xf8, 0xb1, 0xa2, 0xc1, 0x7a,
	0x45, 0x83, 0x2f, 0x0d, 0x25, 0xdf, 0x1b, 0x1a, 0x2c, 0x1b, 0x4a, 0xce, 0x1b, 0x4a, 0x7e, 0x37,
	0x94, 0xfc, 0x6d, 0x68, 0xb0, 0x6e, 0x28, 0xf9, 0xfa, 0x87, 0x06, 0x1f, 0x06, 0xfe, 0xf2, 0x67,
	0x03, 0x77, 0xa1, 0xcf, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xcd, 0x38, 0x3c, 0xba, 0x02,
	0x00, 0x00,
}
