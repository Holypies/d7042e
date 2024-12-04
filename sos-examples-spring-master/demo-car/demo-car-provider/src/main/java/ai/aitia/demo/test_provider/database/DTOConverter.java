package ai.aitia.demo.test_provider.database;

import ai.aitia.demo.test_common.dto.TestResponseDTO;
import ai.aitia.demo.test_provider.entity.Test;
import java.util.ArrayList;
import java.util.List;
import org.springframework.util.Assert;

public class DTOConverter {

    //=================================================================================================
    // methods

    //-------------------------------------------------------------------------------------------------
    public static TestResponseDTO convertTestToTestResponseDTO(
        final Test test
    ) {
        Assert.notNull(test, "Test is null");
        return new TestResponseDTO(test.getId(), test.getIsOn());
    }

    //-------------------------------------------------------------------------------------------------
    public static List<TestResponseDTO> convertCarListToCarResponseDTOList(
        final List<Test> tests
    ) {
        Assert.notNull(tests, "test list is null");
        final List<TestResponseDTO> testResponse = new ArrayList<>(
            tests.size()
        );
        for (final Test test : tests) {
            testResponse.add(convertTestToTestResponseDTO(test));
        }
        return testResponse;
    }

    //=================================================================================================
    // assistant methods

    //-------------------------------------------------------------------------------------------------
    public DTOConverter() {
        throw new UnsupportedOperationException();
    }
}
