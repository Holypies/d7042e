package ai.aitia.demo.test_provider;

import eu.arrowhead.common.CommonConstants;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.ComponentScan;

@SpringBootApplication
@ComponentScan(
    basePackages = {
        CommonConstants.BASE_PACKAGE, TestProviderConstants.BASE_PACKAGE,
    }
)
public class TestProviderMain {

    //=================================================================================================
    // methods

    //-------------------------------------------------------------------------------------------------
    public static void main(final String[] args) {
        SpringApplication.run(TestProviderMain.class, args);
    }
}
