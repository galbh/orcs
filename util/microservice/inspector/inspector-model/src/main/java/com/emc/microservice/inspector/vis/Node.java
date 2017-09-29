/*
 * Copyright (c) 2016 EMC Corporation All Rights Reserved
 */
package com.emc.microservice.inspector.vis;

/**
 * Created by liebea on 1/7/15.
 * Drink responsibly
 */
public class Node {
    private final String id;
    private final String label;
    private final String shape;

    // Required by stupid jackson
    private Node() {
        this(null, null, null);
    }

    public Node(String id, String label, String shape) {
        this.id = id;
        this.label = label;
        this.shape = shape;
    }

    public String getId() {
        return id;
    }

    public String getLabel() {
        return label;
    }

    public String getShape() {
        return shape;
    }
}
