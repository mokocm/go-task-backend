// @generated by protoc-gen-connect-query v0.4.1 with parameter "target=ts"
// @generated from file proto/task/v1/service.proto (package todo.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { createQueryService } from "@bufbuild/connect-query";
import { MethodKind } from "@bufbuild/protobuf";
import { CreateTaskRequest, CreateTaskResponse, GetTaskRequest, GetTaskResponse, GetTasksRequest, GetTasksResponse } from "./service_pb.js";

export const typeName = "todo.v1.TaskService";

/**
 * タスクIDからタスクを取得する
 *
 * @generated from rpc todo.v1.TaskService.GetTask
 */
export const getTask = createQueryService({
  service: {
    methods: {
      getTask: {
        name: "GetTask",
        kind: MethodKind.Unary,
        I: GetTaskRequest,
        O: GetTaskResponse,
      },
    },
    typeName: "todo.v1.TaskService",
  },
}).getTask;

/**
 * タスク一覧を取得する
 *
 * @generated from rpc todo.v1.TaskService.GetTasks
 */
export const getTasks = createQueryService({
  service: {
    methods: {
      getTasks: {
        name: "GetTasks",
        kind: MethodKind.Unary,
        I: GetTasksRequest,
        O: GetTasksResponse,
      },
    },
    typeName: "todo.v1.TaskService",
  },
}).getTasks;

/**
 * タスクを作成する
 *
 * @generated from rpc todo.v1.TaskService.CreateTask
 */
export const createTask = createQueryService({
  service: {
    methods: {
      createTask: {
        name: "CreateTask",
        kind: MethodKind.Unary,
        I: CreateTaskRequest,
        O: CreateTaskResponse,
      },
    },
    typeName: "todo.v1.TaskService",
  },
}).createTask;