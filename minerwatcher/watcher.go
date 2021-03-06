package main

import (
    "github.com/buger/jsonparser"
	"strings"
	"net/http"
	"os/exec"
	"log"
	"fmt"
)

#rigs
rig1 := []array{"4fe11f", 1000}
rig2 := []array{"501b79", 1000}
rigs := []array{rig1, rig2}


data := []byte(`{"rigs":{"501d77":{"condition":"unreachable","version":"1.2.5","driver":"amdgpu","miner":"sgminer-gm-xmr","gpus":"13","miner_instance":0,"miner_hashes":"716.10 755.50 665.40 620.00 00.00 661.20 660.60 620.10 684.80 696.80 696.80 696.40 742.40","bioses":"113-V34121-X3 113-C9403100_101 113-D0090101_100 113-WE353FU-M73 113-C9403100_101 113-WE353GU-M7D 113-WE353GU-M7D 113-WE353FU-M73 113-D0003400_100 113-D0003400_100 113-D0003400_100 113-D0003400_100 113-C9403100_101","meminfo":"GPU0:01.00.0:Radeon RX 470:113-V34121-X3:SK Hynix H5GQ8H24MJR\nGPU1:04.00.0:Radeon RX 580:113-C9403100_101:Samsung K4G80325FB\nGPU2:05.00.0:Radeon RX 480:113-D0090101_100:Samsung K4G80325FB\nGPU3:06.00.0:Radeon RX 470:113-WE353FU-M73:Samsung K4G20325FS\nGPU4:07.00.0:Radeon RX 580:113-C9403100_101:Samsung K4G80325FB\nGPU5:08.00.0:Radeon RX 470:113-WE353GU-M7D:Samsung K4G80325FB\nGPU6:09.00.0:Radeon RX 470:113-WE353GU-M7D:Samsung K4G80325FB\nGPU7:0a.00.0:Radeon RX 470:113-WE353FU-M73:Samsung K4G20325FS\nGPU8:0d.00.0:Radeon RX 570:113-D0003400_100:SK Hynix H5GC4H24AJR\nGPU9:0e.00.0:Radeon RX 570:113-D0003400_100:Elpida EDW4032BABG\nGPU10:0f.00.0:Radeon RX 570:113-D0003400_100:Elpida EDW4032BABG\nGPU11:10.00.0:Radeon RX 570:113-D0003400_100:Elpida EDW4032BABG\nGPU12:13.00.0:Radeon RX 580:113-C9403100_101:Samsung K4G80325FB","vramsize":"8 8 8 4 8 8 8 4 4 4 4 4 8","drive_name":"KingDian S100 16GB 2017061400403","mobo":"H110 Pro BTC ","lan_chip":"Intel Corporation Ethernet Connection (2) I219-V (rev 31)","connected_displays":"","ram":"16","rack_loc":"miner01","ip":"192.168.0.143","server_time":1507929306,"uptime":"401","miner_secs":320,"rx_kbps":"0.00","tx_kbps":"0.00","load":"0.22","cpu_temp":"59","freespace":3.1,"hash":0,"pool":"xmr-us-west1.nanopool.org:14444","temp":"62.00 77.00 68.00 76.00 66.00 55.00 57.00 51.00 76.00 57.00 65.00 69.00 67.00","powertune":"5 5 5 5 5 5 5 5 5 5 5 5 5","voltage":"1.150 1.150 1.175 1.150 1.150 1.150 1.150 1.150 1.150 1.150 1.150 1.150 1.150","watts":null,"fanrpm":"4500 4500 4500 4500 4500 4500 4500 4500 4500 4500 4500 4500 4500","core":"1206 1350 1330 1236 1350 1236 1236 1236 1320 1250 1250 1250 1350","mem":"1960 2000 2100 1975 300 2100 2100 1975 1900 2000 2000 2000 2000"},"501b79":{"condition":"mining","version":"1.2.5","driver":"amdgpu","miner":"sgminer-gm-xmr","gpus":"5","miner_instance":5,"miner_hashes":"619.00 619.30 619.80 782.70 624.30","bioses":"113-D0003400_100 113-D0003400_100 113-C9403100_101 113-C9403100_101 113-D0003400_100","meminfo":"GPU0:01.00.0:Radeon RX 570:113-D0003400_100:Elpida EDW4032BABG\nGPU1:03.00.0:Radeon RX 570:113-D0003400_100:Elpida EDW4032BABG\nGPU2:04.00.0:Radeon RX 580:113-C9403100_101:Samsung K4G80325FB\nGPU3:05.00.0:Radeon RX 580:113-C9403100_101:Samsung K4G80325FB\nGPU4:08.00.0:Radeon RX 570:113-D0003400_100:Elpida EDW4032BABG","vramsize":"4 4 8 8 4","drive_name":"KingDian S100 16GB 2017061400556","mobo":"H110 Pro BTC ","lan_chip":"Intel Corporation Ethernet Connection (2) I219-V (rev 31)","connected_displays":"","ram":"16","rack_loc":"miner02","ip":"192.168.2.2","server_time":1508184249,"uptime":"25157","miner_secs":25099,"rx_kbps":"0.03","tx_kbps":"0.04","load":"0.29","cpu_temp":"36","freespace":3.1,"hash":3265.1,"pool":"xmr-us-east1.nanopool.org:14444","temp":"43.00 46.00 42.00 40.00 48.00","powertune":"5 5 5 5 5","voltage":"1.150 1.143 1.150 1.150 1.150","watts":null,"fanrpm":"3988 3988 3988 3988 3988","core":"1174 1174 1266 1335 1239","mem":"1750 1750 2000 2000 1750"}},"total_hash":3265.1,"alive_gpus":5,"total_gpus":18,"alive_rigs":1,"total_rigs":2,"current_version":"1.2.5","avg_temp":54.44,"capacity":"27.8","per_info":{"sgminer-gm-xmr":{"hash":3265,"per_alive_gpus":5,"per_total_gpus":18,"per_alive_rigs":1,"per_total_rigs":2,"per_hash-gpu":"653.0","per_hash-rig":"3265.0","current_time":1508184293}}}`)

func alert() {
	if out, err := exec.Command(nma.sh HashWatcher LowHash "Hashing for $miner is $hash" 2).Output(); err == nil {
		log.Info("Alert Sent")
	}
}


func main() {
for each i in rigs

var thash int64
if value, err := jsonparser.Get(data, "rigs", i[0], "hash"); err == nil {
thash = value
}

var ihash int64
if value, err := jsonparser.Get(data, "rigs", i[0], "miner_hashes"); err == nil {
ihash = value
}

ihash := strings.Contains(ihash, "00.00")

if thash < i[1] {
	alert()
}

if ihash {
	alert()
}


}