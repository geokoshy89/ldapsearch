package com.geo.LdapSearch.serviceimpl;

import com.geo.LdapSearch.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.ldap.core.AttributesMapper;
import org.springframework.ldap.core.LdapTemplate;
import org.springframework.stereotype.Component;

@Component
public class LdapUserService implements UserService {
    @Autowired
    private LdapTemplate ldapTemplate;
    @Override
    public String getUser(String userName) {
        var val = ldapTemplate
                .search(
                        "ou=People",
                        "cn=" + userName,
                        new PersonAttributesMapper());
        if(!val.isEmpty()){
            return val.get(0).toString();
        }
        return "No user found";
    }
}
