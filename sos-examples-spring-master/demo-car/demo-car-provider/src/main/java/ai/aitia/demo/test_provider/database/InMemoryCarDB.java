package ai.aitia.demo.test_provider.database;

import ai.aitia.demo.test_provider.entity.Test;
import eu.arrowhead.common.exception.InvalidParameterException;
import java.util.List;
import java.util.concurrent.ConcurrentHashMap;
import org.springframework.stereotype.Component;

@Component
public class InMemoryCarDB extends ConcurrentHashMap<Integer, Test> {

    //=================================================================================================
    // members

    private static final long serialVersionUID = -2462387539362748691L;

    private int idCounter = 0;

    //=================================================================================================
    // methods

    //-------------------------------------------------------------------------------------------------
    public Test create(boolean isOn) {
        /*
		if (isOn == null || isOn.isBlank()) {
			throw new InvalidParameterException("brand is null or empty");
		}
		if (color == null || color.isBlank()) {
			throw new InvalidParameterException("color is null or empty");
		}
		*/
        idCounter++;
        this.put(idCounter, new Test(idCounter, isOn));
        return this.get(idCounter);
    }

    //-------------------------------------------------------------------------------------------------
    public List<Test> getAll() {
        return List.copyOf(this.values());
    }

    //-------------------------------------------------------------------------------------------------
    public Test getById(final int id) {
        if (this.containsKey(id)) {
            return this.get(id);
        } else {
            throw new InvalidParameterException("id '" + id + "' not exists");
        }
    }

    //-------------------------------------------------------------------------------------------------
    public Test updateById(final int id, Boolean isOn) {
        if (this.containsKey(id)) {
            final Test test = this.get(id);
            test.setIsOn(isOn);
            this.put(id, test);
            return test;
        } else {
            throw new InvalidParameterException("id '" + id + "' not exists");
        }
    }

    //-------------------------------------------------------------------------------------------------
    public void removeById(final int id) {
        if (this.containsKey(id)) {
            this.remove(id);
        }
    }
}
