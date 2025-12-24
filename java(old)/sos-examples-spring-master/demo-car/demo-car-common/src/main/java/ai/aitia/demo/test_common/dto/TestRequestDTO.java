package ai.aitia.demo.test_common.dto;

import java.io.Serializable;

public class TestRequestDTO implements Serializable {

    //=================================================================================================
    // members

    private static final long serialVersionUID = -5363562707054976998L;

    private boolean isOn;

    //=================================================================================================
    // methods

    //-------------------------------------------------------------------------------------------------
    public TestRequestDTO(boolean isOn) {
        this.isOn = isOn;
    }

    //-------------------------------------------------------------------------------------------------
    public boolean getIsOn() {
        return isOn;
    }

    //-------------------------------------------------------------------------------------------------
    public void setIsOn(boolean isOn) {
        this.isOn = isOn;
    }
}
