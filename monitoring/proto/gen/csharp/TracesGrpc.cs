// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: monitoring/v1/traces.proto
// </auto-generated>
#pragma warning disable 0414, 1591, 8981, 0612
#region Designer generated code

using grpc = global::Grpc.Core;

namespace Smartpcr.Monitoring.Traces.V1 {
  public static partial class TracesService
  {
    static readonly string __ServiceName = "smartpcr.monitoring.traces.v1.TracesService";

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static void __Helper_SerializeMessage(global::Google.Protobuf.IMessage message, grpc::SerializationContext context)
    {
      #if !GRPC_DISABLE_PROTOBUF_BUFFER_SERIALIZATION
      if (message is global::Google.Protobuf.IBufferMessage)
      {
        context.SetPayloadLength(message.CalculateSize());
        global::Google.Protobuf.MessageExtensions.WriteTo(message, context.GetBufferWriter());
        context.Complete();
        return;
      }
      #endif
      context.Complete(global::Google.Protobuf.MessageExtensions.ToByteArray(message));
    }

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static class __Helper_MessageCache<T>
    {
      public static readonly bool IsBufferMessage = global::System.Reflection.IntrospectionExtensions.GetTypeInfo(typeof(global::Google.Protobuf.IBufferMessage)).IsAssignableFrom(typeof(T));
    }

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static T __Helper_DeserializeMessage<T>(grpc::DeserializationContext context, global::Google.Protobuf.MessageParser<T> parser) where T : global::Google.Protobuf.IMessage<T>
    {
      #if !GRPC_DISABLE_PROTOBUF_BUFFER_SERIALIZATION
      if (__Helper_MessageCache<T>.IsBufferMessage)
      {
        return parser.ParseFrom(context.PayloadAsReadOnlySequence());
      }
      #endif
      return parser.ParseFrom(context.PayloadAsNewBuffer());
    }

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Smartpcr.Monitoring.Traces.V1.CreateSpanRequest> __Marshaller_smartpcr_monitoring_traces_v1_CreateSpanRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Smartpcr.Monitoring.Traces.V1.CreateSpanRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Smartpcr.Monitoring.Traces.V1.CreateSpanResponse> __Marshaller_smartpcr_monitoring_traces_v1_CreateSpanResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Smartpcr.Monitoring.Traces.V1.CreateSpanResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Smartpcr.Monitoring.Traces.V1.GetTraceRequest> __Marshaller_smartpcr_monitoring_traces_v1_GetTraceRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Smartpcr.Monitoring.Traces.V1.GetTraceRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Smartpcr.Monitoring.Traces.V1.GetTraceResponse> __Marshaller_smartpcr_monitoring_traces_v1_GetTraceResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Smartpcr.Monitoring.Traces.V1.GetTraceResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Google.Protobuf.WellKnownTypes.Empty> __Marshaller_google_protobuf_Empty = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Google.Protobuf.WellKnownTypes.Empty.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Smartpcr.Monitoring.Traces.V1.Traces> __Marshaller_smartpcr_monitoring_traces_v1_Traces = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Smartpcr.Monitoring.Traces.V1.Traces.Parser));

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Smartpcr.Monitoring.Traces.V1.CreateSpanRequest, global::Smartpcr.Monitoring.Traces.V1.CreateSpanResponse> __Method_CreateSpan = new grpc::Method<global::Smartpcr.Monitoring.Traces.V1.CreateSpanRequest, global::Smartpcr.Monitoring.Traces.V1.CreateSpanResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "CreateSpan",
        __Marshaller_smartpcr_monitoring_traces_v1_CreateSpanRequest,
        __Marshaller_smartpcr_monitoring_traces_v1_CreateSpanResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Smartpcr.Monitoring.Traces.V1.GetTraceRequest, global::Smartpcr.Monitoring.Traces.V1.GetTraceResponse> __Method_GetTrace = new grpc::Method<global::Smartpcr.Monitoring.Traces.V1.GetTraceRequest, global::Smartpcr.Monitoring.Traces.V1.GetTraceResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "GetTrace",
        __Marshaller_smartpcr_monitoring_traces_v1_GetTraceRequest,
        __Marshaller_smartpcr_monitoring_traces_v1_GetTraceResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Google.Protobuf.WellKnownTypes.Empty, global::Smartpcr.Monitoring.Traces.V1.Traces> __Method_GetAllTraces = new grpc::Method<global::Google.Protobuf.WellKnownTypes.Empty, global::Smartpcr.Monitoring.Traces.V1.Traces>(
        grpc::MethodType.Unary,
        __ServiceName,
        "GetAllTraces",
        __Marshaller_google_protobuf_Empty,
        __Marshaller_smartpcr_monitoring_traces_v1_Traces);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::Smartpcr.Monitoring.Traces.V1.TracesReflection.Descriptor.Services[0]; }
    }

    /// <summary>Base class for server-side implementations of TracesService</summary>
    [grpc::BindServiceMethod(typeof(TracesService), "BindService")]
    public abstract partial class TracesServiceBase
    {
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Smartpcr.Monitoring.Traces.V1.CreateSpanResponse> CreateSpan(global::Smartpcr.Monitoring.Traces.V1.CreateSpanRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Smartpcr.Monitoring.Traces.V1.GetTraceResponse> GetTrace(global::Smartpcr.Monitoring.Traces.V1.GetTraceRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Smartpcr.Monitoring.Traces.V1.Traces> GetAllTraces(global::Google.Protobuf.WellKnownTypes.Empty request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

    }

    /// <summary>Client for TracesService</summary>
    public partial class TracesServiceClient : grpc::ClientBase<TracesServiceClient>
    {
      /// <summary>Creates a new client for TracesService</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public TracesServiceClient(grpc::ChannelBase channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for TracesService that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public TracesServiceClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected TracesServiceClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected TracesServiceClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Smartpcr.Monitoring.Traces.V1.CreateSpanResponse CreateSpan(global::Smartpcr.Monitoring.Traces.V1.CreateSpanRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CreateSpan(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Smartpcr.Monitoring.Traces.V1.CreateSpanResponse CreateSpan(global::Smartpcr.Monitoring.Traces.V1.CreateSpanRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_CreateSpan, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Smartpcr.Monitoring.Traces.V1.CreateSpanResponse> CreateSpanAsync(global::Smartpcr.Monitoring.Traces.V1.CreateSpanRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CreateSpanAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Smartpcr.Monitoring.Traces.V1.CreateSpanResponse> CreateSpanAsync(global::Smartpcr.Monitoring.Traces.V1.CreateSpanRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_CreateSpan, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Smartpcr.Monitoring.Traces.V1.GetTraceResponse GetTrace(global::Smartpcr.Monitoring.Traces.V1.GetTraceRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetTrace(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Smartpcr.Monitoring.Traces.V1.GetTraceResponse GetTrace(global::Smartpcr.Monitoring.Traces.V1.GetTraceRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_GetTrace, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Smartpcr.Monitoring.Traces.V1.GetTraceResponse> GetTraceAsync(global::Smartpcr.Monitoring.Traces.V1.GetTraceRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetTraceAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Smartpcr.Monitoring.Traces.V1.GetTraceResponse> GetTraceAsync(global::Smartpcr.Monitoring.Traces.V1.GetTraceRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_GetTrace, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Smartpcr.Monitoring.Traces.V1.Traces GetAllTraces(global::Google.Protobuf.WellKnownTypes.Empty request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetAllTraces(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Smartpcr.Monitoring.Traces.V1.Traces GetAllTraces(global::Google.Protobuf.WellKnownTypes.Empty request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_GetAllTraces, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Smartpcr.Monitoring.Traces.V1.Traces> GetAllTracesAsync(global::Google.Protobuf.WellKnownTypes.Empty request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetAllTracesAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Smartpcr.Monitoring.Traces.V1.Traces> GetAllTracesAsync(global::Google.Protobuf.WellKnownTypes.Empty request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_GetAllTraces, null, options, request);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected override TracesServiceClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new TracesServiceClient(configuration);
      }
    }

    /// <summary>Creates service definition that can be registered with a server</summary>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static grpc::ServerServiceDefinition BindService(TracesServiceBase serviceImpl)
    {
      return grpc::ServerServiceDefinition.CreateBuilder()
          .AddMethod(__Method_CreateSpan, serviceImpl.CreateSpan)
          .AddMethod(__Method_GetTrace, serviceImpl.GetTrace)
          .AddMethod(__Method_GetAllTraces, serviceImpl.GetAllTraces).Build();
    }

    /// <summary>Register service method with a service binder with or without implementation. Useful when customizing the service binding logic.
    /// Note: this method is part of an experimental API that can change or be removed without any prior notice.</summary>
    /// <param name="serviceBinder">Service methods will be bound by calling <c>AddMethod</c> on this object.</param>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static void BindService(grpc::ServiceBinderBase serviceBinder, TracesServiceBase serviceImpl)
    {
      serviceBinder.AddMethod(__Method_CreateSpan, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Smartpcr.Monitoring.Traces.V1.CreateSpanRequest, global::Smartpcr.Monitoring.Traces.V1.CreateSpanResponse>(serviceImpl.CreateSpan));
      serviceBinder.AddMethod(__Method_GetTrace, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Smartpcr.Monitoring.Traces.V1.GetTraceRequest, global::Smartpcr.Monitoring.Traces.V1.GetTraceResponse>(serviceImpl.GetTrace));
      serviceBinder.AddMethod(__Method_GetAllTraces, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Google.Protobuf.WellKnownTypes.Empty, global::Smartpcr.Monitoring.Traces.V1.Traces>(serviceImpl.GetAllTraces));
    }

  }
}
#endregion
