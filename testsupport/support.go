/***** BEGIN LICENSE BLOCK *****
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this file,
# You can obtain one at http://mozilla.org/MPL/2.0/.
#
# The Initial Developer of the Original Code is the Mozilla Foundation.
# Portions created by the Initial Developer are Copyright (C) 2012
# the Initial Developer. All Rights Reserved.
#
# Contributor(s):
#   Rob Miller (rmiller@mozilla.com)
#
# ***** END LICENSE BLOCK *****/
package testsupport

import (
	"log"
)

type SimpleT struct {}

func (*SimpleT) Errorf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (*SimpleT) Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}