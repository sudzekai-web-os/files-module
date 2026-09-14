package statformat

type StatFormat struct {
	Key    string
	Format string
}

var (
	Name = StatFormat{"Name", "%n"}

	// использование данного паттерна не предусмотрено, так как команда stat не имеет специализированного формата
	FullName = StatFormat{"FullName", ""}

	Type       = StatFormat{"Type", "%F"}
	Size       = StatFormat{"Size", "%s"}
	Inode      = StatFormat{"Inode", "%i"}
	Links      = StatFormat{"Links", "%h"}
	DeviceID   = StatFormat{"DeviceID", "%d"}
	RDevice    = StatFormat{"RDevice", "%r"}
	BlockSize  = StatFormat{"BlockSize", "%B"}
	Blocks     = StatFormat{"Blocks", "%b"}
	Owner      = StatFormat{"Owner", "%U"}
	OwnerID    = StatFormat{"OwnerID", "%u"}
	Group      = StatFormat{"Group", "%G"}
	GroupID    = StatFormat{"GroupID", "%g"}
	LinkTarget = StatFormat{"LinkTarget", "%N"}

	Permissions = StatFormat{"Permissions", "%a"}

	// использование данного паттерна не предусмотрено, так как команда stat не имеет специализированного формата
	SpecialPermissions = StatFormat{"SpecialPermissions", ""}
	// использование данного паттерна не предусмотрено, так как команда stat не имеет специализированного формата
	UserPermissions = StatFormat{"UserPermissions", ""}
	// использование данного паттерна не предусмотрено, так как команда stat не имеет специализированного формата
	GroupPermissions = StatFormat{"GroupPermissions", ""}
	// использование данного паттерна не предусмотрено, так как команда stat не имеет специализированного формата
	OthersPermissions = StatFormat{"OthersPermissions", ""}

	BirthDateTime        = StatFormat{"BirthDateTime", "%W"}
	ModificationDateTime = StatFormat{"ModificationDateTime", "%Y"}
	ChangeDateTime       = StatFormat{"ChangeDateTime", "%Z"}
	AccessDateTime       = StatFormat{"AccessDateTime", "%X"}
)
