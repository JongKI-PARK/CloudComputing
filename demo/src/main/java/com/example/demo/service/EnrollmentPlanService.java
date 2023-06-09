package com.example.demo.service;

import com.example.demo.repository.EnrollmentPlanRepository;
import com.example.demo.service.dto.EnrollmentCreateRequestDto;
import com.example.demo.service.dto.EnrollmentResponseDto;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
@Transactional(readOnly = true)
@RequiredArgsConstructor
public class EnrollmentPlanService {
    private final EnrollmentPlanRepository repository;

    @Transactional
    public Long saveEnrollmentPlan(EnrollmentCreateRequestDto requestDto) {
        return repository.save(requestDto.convertToEntity()).getId();
    }

    public List<EnrollmentResponseDto> getStudentEnrollmentPlan(Long studentId) {
        //get EnrollmentPlan List
        //convert EnrollmentPlan to EnrollmentResponseDto
        return repository.findByStudentId(studentId).stream()
                .map(EnrollmentResponseDto::new)
                .toList();
    }
}
