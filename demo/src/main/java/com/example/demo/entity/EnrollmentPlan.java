package com.example.demo.entity;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Id;
import lombok.Getter;
import lombok.NoArgsConstructor;

@Entity
@Getter
@NoArgsConstructor
public class EnrollmentPlan {
    @Id @GeneratedValue
    private Long id;

    @Column(name = "student_id")
    private Long studentId;

    @Column(name = "subject_id")
    private Long subjectId;

    public EnrollmentPlan(Long studentId, Long subjectId) {
        this.studentId = studentId;
        this.subjectId = subjectId;
    }
}
