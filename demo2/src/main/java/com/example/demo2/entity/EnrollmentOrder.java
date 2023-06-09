package com.example.demo2.entity;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;

@Entity
@Getter
@NoArgsConstructor
@AllArgsConstructor
public class EnrollmentOrder {
    @Id
    @Column(name = "subject_id")
    private Long subjectId;

    @Column(name = "enrollment_cap")
    private int enrollmentCap;

    @Column(name = "enrollment_current")
    private int enrollmentCurrent;

    public void addEnrollmentNumber() {
        enrollmentCurrent += 1;
    }
}

