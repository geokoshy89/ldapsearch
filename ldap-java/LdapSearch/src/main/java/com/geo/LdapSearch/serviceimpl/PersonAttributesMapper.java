package com.geo.LdapSearch.serviceimpl;

import org.springframework.ldap.NamingException;
import org.springframework.ldap.core.AttributesMapper;

import javax.naming.directory.Attributes;

public class PersonAttributesMapper implements AttributesMapper<Person> {
    @Override
    public Person mapFromAttributes(Attributes attrs) throws NamingException, javax.naming.NamingException {
        Person person = new Person();
        person.setFullName((String) attrs.get("cn").get());
        person.setLastName((String) attrs.get("sn").get());
        return person;
    }
}
