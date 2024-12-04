package ai.aitia.demo.test_provider.controller;

import ai.aitia.demo.test_common.dto.TestRequestDTO;
import ai.aitia.demo.test_common.dto.TestResponseDTO;
import ai.aitia.demo.test_provider.TestProviderConstants;
import ai.aitia.demo.test_provider.database.DTOConverter;
import ai.aitia.demo.test_provider.database.InMemoryTestDB;
import ai.aitia.demo.test_provider.entity.Test;
import eu.arrowhead.common.exception.BadPayloadException;
import java.util.ArrayList;
import java.util.List;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping(TestProviderConstants.TEST_URI)
public class TestServiceController {

    //=================================================================================================
    // members

    @Autowired
    private InMemoryTestDB testDB;

    //=================================================================================================
    // methods

    //-------------------------------------------------------------------------------------------------
    @GetMapping(produces = MediaType.APPLICATION_JSON_VALUE)
    @ResponseBody
    public List<TestResponseDTO> getTests(
        @RequestParam(
            name = TestProviderConstants.REQUEST_PARAM_ISON,
            required = false
        ) final Boolean isOn
    ) {
        final List<TestResponseDTO> response = new ArrayList<>();
        for (final Test test : testDB.getAll()) {
            /*
            boolean toAdd = true;

            if (!test.getIsOn().equalsIgnoreCase(brand)) {
                toAdd = false;
            }
            */
            //if (toAdd) {
            response.add(DTOConverter.convertTestToTestResponseDTO(test));
            //}
        }
        return response;
    }

    //-------------------------------------------------------------------------------------------------
    @GetMapping(
        path = TestProviderConstants.BY_ID_PATH,
        produces = MediaType.APPLICATION_JSON_VALUE
    )
    @ResponseBody
    public TestResponseDTO getTestById(
        @PathVariable(
            value = TestProviderConstants.PATH_VARIABLE_ID
        ) final int id
    ) {
        return DTOConverter.convertTestToTestResponseDTO(testDB.getById(id));
    }

    //-------------------------------------------------------------------------------------------------
    @PostMapping(
        consumes = MediaType.APPLICATION_JSON_VALUE,
        produces = MediaType.APPLICATION_JSON_VALUE
    )
    @ResponseBody
    public TestResponseDTO createTest(@RequestBody final TestRequestDTO dto) {
        /*
        if (dto.getBrand() == null || dto.getBrand().isBlank()) {
            throw new BadPayloadException("brand is null or blank");
        }
        if (dto.getColor() == null || dto.getColor().isBlank()) {
            throw new BadPayloadException("color is null or blank");
        }
        */
        final Test test = testDB.create(dto.getIsOn());
        return DTOConverter.convertTestToTestResponseDTO(test);
    }

    //-------------------------------------------------------------------------------------------------
    @PutMapping(
        path = TestProviderConstants.BY_ID_PATH,
        consumes = MediaType.APPLICATION_JSON_VALUE,
        produces = MediaType.APPLICATION_JSON_VALUE
    )
    @ResponseBody
    public TestResponseDTO updateTest(
        @PathVariable(
            name = TestProviderConstants.PATH_VARIABLE_ID
        ) final int id,
        @RequestBody final TestRequestDTO dto
    ) {
        /*
        if (dto.getBrand() == null || dto.getBrand().isBlank()) {
            throw new BadPayloadException("brand is null or blank");
        }
        if (dto.getColor() == null || dto.getColor().isBlank()) {
            throw new BadPayloadException("color is null or blank");
        }
        */
        final Test test = testDB.updateById(id, dto.getIsOn());
        return DTOConverter.convertTestToTestResponseDTO(test);
    }

    //-------------------------------------------------------------------------------------------------
    @DeleteMapping(path = TestProviderConstants.BY_ID_PATH)
    public void removeTestById(
        @PathVariable(
            value = TestProviderConstants.PATH_VARIABLE_ID
        ) final int id
    ) {
        testDB.removeById(id);
    }
}
