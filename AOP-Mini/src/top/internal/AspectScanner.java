package top.internal;

import top.annotations.After;
import top.aspects.AspectRegistry;

import java.io.File;
import java.lang.reflect.Method;
import java.net.URL;
import java.util.ArrayList;
import java.util.List;


public class AspectScanner {

    private List<Object> registeredAspects;
    private AspectRegistry aspectRegistry;

    public AspectScanner() {
        this.registeredAspects = new ArrayList<>();
        this.aspectRegistry = new AspectRegistry();
    }

    public void registerAspect(Object aspect) {
        this.registeredAspects.add(aspect);
    }

    public void scanCodeForAspects() {
        String packageName = "top";

        try {
            String packagePath = packageName.replace('.', '/');
            ClassLoader classLoader = Thread.currentThread().getContextClassLoader();
            URL packageURL = classLoader.getResource(packagePath);

            if (packageURL == null) {
                System.out.println("Package not found: " + packageName);
                return;
            }

            File directory = new File(packageURL.toURI());
            List<File> classFiles = new ArrayList<>();
            listClassFiles(directory, classFiles);

            for (File classFile : classFiles) {
                // Convert file path back to class name
                String path = classFile.getAbsolutePath();
                String basePath = directory.getAbsolutePath();

                String className = path.substring(basePath.length() + 1, path.length() - 6)  // remove ".class"
                        .replace(File.separatorChar, '.');

                // Load the class
                Class<?> clazz = Class.forName(packageName + "." + className);

                // Now inspect methods
                System.out.println("Class: " + clazz.getName());
                for (Method method : clazz.getDeclaredMethods()) {
                    if (method.isAnnotationPresent(After.class)) {
                        aspectRegistry.add(After.class.getName(), method);
                        System.out.println(" Registered After Aspect Method: " + method.getName());
                    }
                }
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    private static void listClassFiles(File directory, List<File> classFiles) {
        for (File file : directory.listFiles()) {
            if (file.isDirectory()) {
                listClassFiles(file, classFiles);
            } else if (file.getName().endsWith(".class")) {
                classFiles.add(file);
            }
        }
    }
}
