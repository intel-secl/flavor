package flavors

/**
 * 
 * @author purvades
 */

type Description struct {
    FlavorPart      string      `json:"flavor_part"`
    Source          string      `json:"source"`
    Label           string      `json:"label"`
    IPAddress       string      `json:"ip_address"`
    BiosName        string      `json:"bios_name"`
    BiosVersion     string      `json:"bios_version"`
    OSName          string      `json:"os_name"`
    OSVersion       string      `json:"os_version"`
    VMMName         string      `json:"vmm_name"`
    VMMVersion      string      `json:"vmm_version"`
    TPMVersion      string      `json:"tpm_version"`
    HardwareUUID    string      `json:"hardware_uuid"`
    Comment         string      `json:"comment"`
}