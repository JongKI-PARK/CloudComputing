package com.example.demo2.service;

import com.example.demo2.entity.EnrollmentComplete;
import com.example.demo2.entity.EnrollmentOrder;
import com.example.demo2.repository.EnrollmentCompleteRepository;
import com.example.demo2.repository.EnrollmentOrderRepository;
import com.example.demo2.service.dto.EnrollmentCompleteDto;
import com.example.demo2.service.dto.EnrollmentOrderRequest;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
@Transactional(readOnly = true)
@RequiredArgsConstructor
public class EnrollmentProcess {
    private final EnrollmentOrderRepository orderRepository;
    private final EnrollmentCompleteRepository completeRepository;

    @Transactional
    public EnrollmentCompleteDto saveEnrollmentComplete(EnrollmentOrderRequest request) {
        EnrollmentOrder order = orderRepository.findById(request.getSubjectId())
                .orElseThrow(RuntimeException::new);

        if(!(order.getEnrollmentCurrent() < order.getEnrollmentCap()))
            return null;

        //update
        order.addEnrollmentNumber();

        EnrollmentComplete savedComplete = completeRepository.save(request.convertToEnrollmentComplete());
        return new EnrollmentCompleteDto(savedComplete.getStudentId(), savedComplete.getSubjectId());
    }
}
