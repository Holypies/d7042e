package ai.aitia.demo.test_provider;

public class TestProviderConstants {

    //=================================================================================================
    // members

    public static final String BASE_PACKAGE = "ai.aitia";

    public static final String CREATE_TEST_SERVICE_DEFINITION = "create-test";
    public static final String GET_TEST_SERVICE_DEFINITION = "get-test";
    public static final String INTERFACE_SECURE = "HTTP-SECURE-JSON";
    public static final String INTERFACE_INSECURE = "HTTP-INSECURE-JSON";
    public static final String HTTP_METHOD = "http-method";
    public static final String TEST_URI = "/test";
    public static final String BY_ID_PATH = "/{id}";
    public static final String PATH_VARIABLE_ID = "id";
    public static final String REQUEST_PARAM_KEY_ISON = "request-param-isOn";
    public static final String REQUEST_PARAM_ISON = "isOn";

    //=================================================================================================
    // assistant methods

    //-------------------------------------------------------------------------------------------------
    private TestProviderConstants() {
        throw new UnsupportedOperationException();
    }
}
