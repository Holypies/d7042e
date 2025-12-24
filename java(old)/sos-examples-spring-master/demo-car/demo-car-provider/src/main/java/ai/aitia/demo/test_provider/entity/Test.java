package ai.aitia.demo.test_provider.entity;

public class Test {

    //=================================================================================================
    // members

    private final int id;
    private boolean isOn;

    //=================================================================================================
    // methods

    //-------------------------------------------------------------------------------------------------
    public Test(final int id, boolean isOn) {
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
    public void setIsOn(boolean isOn) {
        this.isOn = isOn;
    }
}
