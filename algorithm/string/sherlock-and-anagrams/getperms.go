package main

func get_perms(str string) PERMS {
    perms := make(PERMS)
    strlen := len(str)
    tmp:=0
    ptrn:="%0"+fmt.Sprintf("%v", strlen)+"b"
    for i:=0; tmp<(1<<uint(strlen))-1; i++ {
        tmp += 1
        permcode := fmt.Sprintf(ptrn, tmp)
        var perm string
        for idx, v := range permcode {
            if v == '1' { perm += string(str[idx]) }
        }
        perms[len(perm)] = append(perms[len(perm)], perm)
    }
    return perms
}
