package com.example.aop.dao;

import com.example.aop.entity.Employee;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.RowMapper;
import org.springframework.stereotype.Repository;

import javax.sql.DataSource;
import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;

@Slf4j
@Repository
public class EmployeeDao {

    private JdbcTemplate jdbcTemplate;

    @Autowired
    public void setDataSource(DataSource dataSource){
        this.jdbcTemplate = new JdbcTemplate(dataSource);
    }

    public int insert(Employee employee){
        int n = this.jdbcTemplate.update("insert into employees values (?,?,?)", employee.getId(), employee.getName(), employee.getDeptNo());
        return n;
    }
    public int delete(Employee employee){
        int n = this.jdbcTemplate.update("delete from employees where id = ?", employee.getId());
        return n;
    }
    public Employee select(int id){
        return this.jdbcTemplate.queryForObject(
                "select * from employees where id = ?",
                new RowMapper<Employee>() {
                    public Employee mapRow(ResultSet rs, int rowNum) throws SQLException {
                        Employee employee = new Employee();
                        employee.setId(rs.getInt(1));
                        employee.setName(rs.getString(2));
                        employee.setDeptNo(rs.getInt(3));
                        return employee;
                    }
                },
                id
                );
    }
}
