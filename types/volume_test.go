package types

import (
	"reflect"
	"testing"

	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/ghodss/yaml"
	"github.com/kr/pretty"

	"github.com/koki/short/util"
)

var kokiHostPath0 = Volume{
	HostPath: &HostPathVolume{
		Path: "/path",
		Type: HostPathDirectoryOrCreate,
	},
}

var sizeLimit0 = resource.MustParse("10Gi")
var kokiEmptyDir0 = Volume{
	EmptyDir: &EmptyDirVolume{
		Medium:    "memory",
		SizeLimit: &sizeLimit0,
	},
}
var kokiEmptyDir1 = Volume{
	EmptyDir: &EmptyDirVolume{},
}

var kokiGcePD0 = Volume{
	GcePD: &GcePDVolume{
		PDName:    "data-disk",
		FSType:    "ext4",
		Partition: 1,
		ReadOnly:  true,
	},
}
var kokiGcePD1 = Volume{
	GcePD: &GcePDVolume{
		PDName: "data-disk",
	},
}

var kokiAwsEBS0 = Volume{
	AwsEBS: &AwsEBSVolume{
		VolumeID:  "volume-id",
		FSType:    "ext4",
		Partition: 1,
		ReadOnly:  true,
	},
}
var kokiAwsEBS1 = Volume{
	AwsEBS: &AwsEBSVolume{
		VolumeID: "volume-id",
	},
}

var azureDiskCachingMode0 = AzureDataDiskCachingReadWrite
var azureDiskKind0 = AzureSharedBlobDisk
var kokiAzureDisk0 = Volume{
	AzureDisk: &AzureDiskVolume{
		DiskName:    "test.vhd",
		DataDiskURI: "https://someaccount.blob.microsoft.net/vhds/test.vhd",
		CachingMode: &azureDiskCachingMode0,
		FSType:      "ext4",
		ReadOnly:    true,
		Kind:        &azureDiskKind0,
	},
}

var kokiAzureFile0 = Volume{
	AzureFile: &AzureFileVolume{
		SecretName: "azure-secret",
		ShareName:  "k8stest",
		ReadOnly:   true,
	},
}
var kokiAzureFile1 = Volume{
	AzureFile: &AzureFileVolume{
		SecretName: "azure-secret",
		ShareName:  "k8stest",
	},
}

var kokiCephFS0 = Volume{
	CephFS: &CephFSVolume{
		Monitors: []string{
			"1.2.3.4:6789",
			"1.2.3.5:6789",
		},
		Path: "/path",
		User: "admin",
		SecretFileOrRef: &CephFSSecretFileOrRef{
			File: "/etc/ceph/admin.secret",
		},
		ReadOnly: true,
	},
}
var kokiCephFS1 = Volume{
	CephFS: &CephFSVolume{
		Monitors: []string{
			"1.2.3.4:6789",
			"1.2.3.5:6789",
		},
		Path: "/path",
		User: "admin",
		SecretFileOrRef: &CephFSSecretFileOrRef{
			Ref: "secret-name",
		},
		ReadOnly: true,
	},
}

var kokiCinder0 = Volume{
	Cinder: &CinderVolume{
		VolumeID: "bd82f7e2-wece-4c01-a505-4acf60b07f4a",
		FSType:   "ext4",
		ReadOnly: true,
	},
}
var kokiCinder1 = Volume{
	Cinder: &CinderVolume{
		VolumeID: "bd82f7e2-wece-4c01-a505-4acf60b07f4a",
	},
}

var kokiFibreChannel0 = Volume{
	FibreChannel: &FibreChannelVolume{
		TargetWWNs: []string{
			"500a0982991b8dc5",
			"500a0982891b8dc5",
		},
		Lun:      util.Int32Ptr(2),
		FSType:   "ext4",
		ReadOnly: true,
		WWIDs: []string{
			"actually, WWIDs should not be specified here",
			"either wwids or (targetwwns + lun), not both",
		},
	},
}

var kokiFlexVolume0 = Volume{
	Flex: &FlexVolume{
		Driver:    "kubernetes.io/lvm",
		FSType:    "ext4",
		SecretRef: "secret-name",
		ReadOnly:  true,
		Options: map[string]string{
			"volumeID":    "vol1",
			"size":        "1000m",
			"volumegroup": "kube_vg",
		},
	},
}
var kokiFlexVolume1 = Volume{
	Flex: &FlexVolume{
		Driver: "kubernetes.io/lvm",
	},
}

var kokiFlockerVolume0 = Volume{
	Flocker: &FlockerVolume{
		DatasetUUID: "flocker_uuid",
	},
}

var kokiGlusterfsVolume0 = Volume{
	Glusterfs: &GlusterfsVolume{
		EndpointsName: "glusterfs-cluster",
		Path:          "kube_vol",
		ReadOnly:      true,
	},
}

var kokiISCSIVolume0 = Volume{
	ISCSI: &ISCSIVolume{
		TargetPortal:   "1.2.3.4:3260",
		IQN:            "iqn.2015-02.example.com:test",
		Lun:            0,
		ISCSIInterface: "default",
		FSType:         "ext4",
		ReadOnly:       true,
		Portals: []string{
			"1.2.3.5:3260",
			"1.2.3.6:3260",
		},
		DiscoveryCHAPAuth: true,
		SessionCHAPAuth:   true,
		SecretRef:         "secret-name",
		InitiatorName:     "iqn.1996-04.de.suse:linux-host1",
	},
}

var kokiNFSVolume0 = Volume{
	NFS: &NFSVolume{
		Server:   "hostname",
		Path:     "/path",
		ReadOnly: true,
	},
}
var kokiNFSVolume1 = Volume{
	NFS: &NFSVolume{
		Server: "1.2.3.4",
		Path:   "/path",
	},
}

var kokiPhotonPDVolume0 = Volume{
	PhotonPD: &PhotonPDVolume{
		PdID:   "some-pdid",
		FSType: "ext4",
	},
}
var kokiPhotonPDVolume1 = Volume{
	PhotonPD: &PhotonPDVolume{
		PdID: "some-pdid",
	},
}

var kokiPortworxVolume0 = Volume{
	Portworx: &PortworxVolume{
		VolumeID: "volume-id",
		FSType:   "ext4",
		ReadOnly: true,
	},
}
var kokiPortworxVolume1 = Volume{
	Portworx: &PortworxVolume{
		VolumeID: "volume-id",
	},
}

var kokiPVCVolume0 = Volume{
	PVC: &PVCVolume{
		ClaimName: "claim-0",
		ReadOnly:  true,
	},
}
var kokiPVCVolume1 = Volume{
	PVC: &PVCVolume{
		ClaimName: "claim-0",
	},
}

var kokiQuobyteVolume0 = Volume{
	Quobyte: &QuobyteVolume{
		Registry: "registry:6789",
		Volume:   "testVolume",
		ReadOnly: true,
		User:     "root",
		Group:    "root",
	},
}

var kokiScaleIOVolume0 = Volume{
	ScaleIO: &ScaleIOVolume{
		Gateway:          "https://localhost:443/api",
		System:           "scaleio",
		SecretRef:        "secret-name",
		SSLEnabled:       true,
		ProtectionDomain: "pd01",
		StoragePool:      "sp01",
		StorageMode:      "ThickProvisioned",
		VolumeName:       "vol-0",
		FSType:           "xfs",
		ReadOnly:         true,
	},
}

func TestVolume(t *testing.T) {
	testVolumeSource(kokiHostPath0, t, true)
	testVolumeSource(kokiEmptyDir0, t, false)
	testVolumeSource(kokiEmptyDir1, t, true)
	testVolumeSource(kokiGcePD0, t, false)
	testVolumeSource(kokiGcePD1, t, true)
	testVolumeSource(kokiAwsEBS0, t, false)
	testVolumeSource(kokiAwsEBS1, t, true)
	testVolumeSource(kokiAzureDisk0, t, false)
	testVolumeSource(kokiAzureFile0, t, true)
	testVolumeSource(kokiAzureFile1, t, true)
	testVolumeSource(kokiCephFS0, t, false)
	testVolumeSource(kokiCephFS1, t, false)
	testVolumeSource(kokiCinder0, t, false)
	testVolumeSource(kokiCinder1, t, true)
	testVolumeSource(kokiFibreChannel0, t, false)
	testVolumeSource(kokiFlexVolume0, t, false)
	testVolumeSource(kokiFlexVolume1, t, true)
	testVolumeSource(kokiFlockerVolume0, t, true)
	testVolumeSource(kokiGlusterfsVolume0, t, false)
	testVolumeSource(kokiISCSIVolume0, t, false)
	testVolumeSource(kokiNFSVolume0, t, true)
	testVolumeSource(kokiNFSVolume1, t, true)
	testVolumeSource(kokiPhotonPDVolume0, t, true)
	testVolumeSource(kokiPhotonPDVolume1, t, true)
	testVolumeSource(kokiPortworxVolume0, t, false)
	testVolumeSource(kokiPortworxVolume1, t, true)
	testVolumeSource(kokiPVCVolume0, t, true)
	testVolumeSource(kokiPVCVolume1, t, true)
	testVolumeSource(kokiQuobyteVolume0, t, false)
	testVolumeSource(kokiScaleIOVolume0, t, false)
}

func isString(data []byte, t *testing.T) bool {
	str := ""
	err := yaml.Unmarshal(data, &str)
	return err == nil
}

func testVolumeSource(kokiVolume Volume, t *testing.T, expectString bool) {
	b, err := yaml.Marshal(kokiVolume)
	if err != nil {
		t.Error(pretty.Sprintf("%s\n%# v", err.Error(), kokiVolume))
		return
	}

	if expectString != isString(b, t) {
		t.Error(pretty.Sprintf("unexpected serialization\n(%s)\n(%# v)", string(b), kokiVolume))
		return
	}

	newVolume := Volume{}

	err = yaml.Unmarshal(b, &newVolume)
	if err != nil {
		t.Error(pretty.Sprintf("%s\n(%s)\n(%# v)", err.Error(), string(b), kokiVolume))
		return
	}

	newB, err := yaml.Marshal(newVolume)
	if err != nil {
		t.Error(pretty.Sprintf("%s\n(%# v)\n(%# v)\n(%s)", err.Error(), newVolume, kokiVolume, string(b)))
		return
	}

	if !reflect.DeepEqual(kokiVolume, newVolume) {
		t.Error(pretty.Sprintf("failed round-trip\n(%# v)\n(%# v)\n(%s)\n(%s)", kokiVolume, newVolume, string(b), string(newB)))
		return
	}
}
