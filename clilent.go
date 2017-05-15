package gobianli



//Every node in LinkedList
type IHost struct{
    host string
    next *IHost
}

//LinkedList
type LinkedList struct{
    Head *IHost
}


