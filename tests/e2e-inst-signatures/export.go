package main

import (
	"github.com/khulnasoft/tracee/tests/e2e-inst-signatures/datasourcetest"
	"github.com/khulnasoft/tracee/types/detect"
)

var ExportedSignatures = []detect.Signature{
	// Instrumentation e2e signatures
	&e2eVfsWrite{},
	&e2eFileModification{},
	&e2eSecurityInodeRename{},
	&e2eContainersDataSource{},
	&e2eBpfAttach{},
	&e2eProcessTreeDataSource{},
	&e2eHookedSyscall{},
	&e2eSignatureDerivation{},
	&e2eDnsDataSource{},
	&e2eWritableDatasourceSig{},
}

var ExportedDataSources = []detect.DataSource{
	datasourcetest.New(),
}
