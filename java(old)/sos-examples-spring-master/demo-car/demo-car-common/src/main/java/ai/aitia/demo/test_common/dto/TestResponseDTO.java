package ai.aitia.demo.test_common.dto;

import java.io.Serializable;

public class TestResponseDTO implements Serializable {

    //=================================================================================================
    // members

    private static final long serialVersionUID = -8371510478751740542L;

    private int id;
    private boolean isOn;

    //=================================================================================================
    // methods

    //-------------------------------------------------------------------------------------------------
    public TestResponseDTO() {}

    //-------------------------------------------------------------------------------------------------
    public TestResponseDTO(final int id, boolean isOn) {
        this.id = id;
        this.isOn = isOn;
    }

    //-------------------------------------------------------------------------------------------------
    public int getId() {
        return id;
    }

    public boolean getIsOn() {
        return isOn;
    }

    //-------------------------------------------------------------------------------------------------
    public void setId(final int id) {
        this.id = id;
    }

    public void setBrand(boolean isOn) {
        this.isOn = isOn;
    }
}
