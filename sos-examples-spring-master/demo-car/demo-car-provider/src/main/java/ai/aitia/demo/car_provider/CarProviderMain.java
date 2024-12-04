package ai.aitia.demo.car_provider;

import eu.arrowhead.common.CommonConstants;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.ComponentScan;

@SpringBootApplication
@ComponentScan(
    basePackages = {
        CommonConstants.BASE_PACKAGE, CarProviderConstants.BASE_PACKAGE,
    }
)
public class CarProviderMain {

    //=================================================================================================
    // methods

    //-------------------------------------------------------------------------------------------------
    public static void main(final String[] args) {
        SpringApplication.run(CarProviderMain.class, args);
    }
}
