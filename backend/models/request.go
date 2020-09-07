package models

type Request struct {
	Raw   []byte        	 `bson:"raw"`
	Host  string		     `bson:"host"`
	Port  int				 `bson:"port"`
	Protocol string        `bson:"protocol"`
}
