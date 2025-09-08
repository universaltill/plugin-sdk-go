package main

import (
  "context"
  "log"
  "google.golang.org/grpc"
  pb "github.com/universaltill/addon-sdk-go/proto/payments"
  "github.com/universaltill/addon-sdk-go/sdk"
)

type server struct{ pb.UnimplementedPaymentsServer }

func (s *server) Metadata(ctx context.Context, _ *pb.Empty) (*pb.AddonMeta, error) {
  return &pb.AddonMeta{Name: "QR Demo PSP", Version: "0.0.1", Methods: []string{"demo_qr"}, PciOutOfScope: true, Regions: []string{"GLOBAL"}}, nil
}
func (s *server) CreateIntent(ctx context.Context, r *pb.CreateIntentRequest) (*pb.CreateIntentResponse, error) {
  return &pb.CreateIntentResponse{IntentId: "demo_1", ProviderRef: "prov_demo"}, nil
}
func (s *server) StartPayment(ctx context.Context, r *pb.StartPaymentRequest) (*pb.StartPaymentResponse, error) {
  return &pb.StartPaymentResponse{Mode: pb.StartPaymentResponse_QR, RequiresPolling: true}, nil
}
func (s *server) GetStatus(ctx context.Context, r *pb.GetStatusRequest) (*pb.GetStatusResponse, error) {
  return &pb.GetStatusResponse{State: pb.GetStatusResponse_CAPTURED, ReceiptText: "Paid (demo)"}, nil
}
func (s *server) CancelIntent(ctx context.Context, r *pb.CancelIntentRequest) (*pb.CancelIntentResponse, error) {
  return &pb.CancelIntentResponse{}, nil
}
func (s *server) Refund(ctx context.Context, r *pb.RefundRequest) (*pb.RefundResponse, error) {
  return &pb.RefundResponse{RefundId: "refund_demo"}, nil
}

func main() {
  lis, err := sdk.Listen(":7001")
  if err != nil { log.Fatal(err) }
  gs := grpc.NewServer()
  pb.RegisterPaymentsServer(gs, &server{})
  log.Println("QR demo PSP listening on :7001")
  log.Fatal(gs.Serve(lis))
}
