// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package fips_state_codes

import (
	"fmt"
)

var stateCodeToVIPSCodeMap = map[string]string{
	"AL": "01",
	"AK": "02",
	"AZ": "04",
	"AR": "05",
	"CA": "06",
	"CO": "08",
	"CT": "09",
	"DE": "10",
	"FL": "12",
	"GA": "13",
	"HI": "15",
	"ID": "16",
	"IL": "17",
	"IN": "18",
	"IA": "19",
	"KS": "20",
	"KY": "21",
	"LA": "22",
	"ME": "23",
	"MD": "24",
	"MA": "25",
	"MI": "26",
	"MN": "27",
	"MS": "28",
	"MO": "29",
	"MT": "30",
	"NE": "31",
	"NV": "32",
	"NH": "33",
	"NJ": "34",
	"NM": "35",
	"NY": "36",
	"NC": "37",
	"ND": "38",
	"OH": "39",
	"OK": "40",
	"OR": "41",
	"PA": "42",
	"RI": "44",
	"SC": "45",
	"SD": "46",
	"TN": "47",
	"TX": "48",
	"UT": "49",
	"VT": "50",
	"VA": "51",
	"WA": "53",
	"WV": "54",
	"WI": "55",
	"WY": "56",
	"AS": "60",
	"GU": "66",
	"MP": "69",
	"PR": "72",
	"VI": "78",
}

func FIPSCodeFromStateCode(stateCode string) (string, error) {
	got, ok := stateCodeToVIPSCodeMap[stateCode]
	if !ok {
		return "", fmt.Errorf("unknown state code: %s", stateCode)
	}
	return got, nil
}

var vipsCodeToStateCodeMap = map[string]string{
	"01": "AL",
	"02": "AK",
	"04": "AZ",
	"05": "AR",
	"06": "CA",
	"08": "CO",
	"09": "CT",
	"10": "DE",
	"12": "FL",
	"13": "GA",
	"15": "HI",
	"16": "ID",
	"17": "IL",
	"18": "IN",
	"19": "IA",
	"20": "KS",
	"21": "KY",
	"22": "LA",
	"23": "ME",
	"24": "MD",
	"25": "MA",
	"26": "MI",
	"27": "MN",
	"28": "MS",
	"29": "MO",
	"30": "MT",
	"31": "NE",
	"32": "NV",
	"33": "NH",
	"34": "NJ",
	"35": "NM",
	"36": "NY",
	"37": "NC",
	"38": "ND",
	"39": "OH",
	"40": "OK",
	"41": "OR",
	"42": "PA",
	"44": "RI",
	"45": "SC",
	"46": "SD",
	"47": "TN",
	"48": "TX",
	"49": "UT",
	"50": "VT",
	"51": "VA",
	"53": "WA",
	"54": "WV",
	"55": "WI",
	"56": "WY",
	"60": "AS",
	"66": "GU",
	"69": "MP",
	"72": "PR",
	"78": "VI",
}

func StateCodeFromFIPSCode(fipsCode string) (string, error) {
	got, ok := vipsCodeToStateCodeMap[fipsCode]
	if !ok {
		return "", fmt.Errorf("unknown fips code: %s", fipsCode)
	}
	return got, nil
}
