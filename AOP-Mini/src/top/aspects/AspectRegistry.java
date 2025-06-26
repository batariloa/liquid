package top.aspects;

import java.lang.reflect.Method;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class AspectRegistry {

    public AspectRegistry() {
        afterAdvice = new HashMap<>();
    }

    private final Map<String, List<Method>> afterAdvice;

    public void add(String aspectName, Method method) {
        afterAdvice.computeIfAbsent(aspectName, k -> new ArrayList<>()).add(method);
    }
}
