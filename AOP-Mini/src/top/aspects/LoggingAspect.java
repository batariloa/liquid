package top.aspects;

import top.annotations.After;
import top.annotations.Aspect;

@Aspect
public class LoggingAspect {

    @After("lol")
    public void logAfter() {
        System.out.println("Well well well ");
    }
}
