package billing

import (
	"context"

	"github.com/arangodb/go-driver"
	pb "github.com/slntopp/nocloud-proto/billing"
	"github.com/slntopp/nocloud/pkg/graph"
	"go.uber.org/zap"
)

type CurrencyServiceServer struct {
	pb.UnimplementedCurrencyServiceServer
	log *zap.Logger

	ctrl graph.CurrencyController

	db driver.Database
}

func NewCurrencyServiceServer(log *zap.Logger, db driver.Database) *CurrencyServiceServer {
	return &CurrencyServiceServer{
		log:  log.Named("CurrencyServer"),
		db:   db,
		ctrl: graph.NewCurrencyController(log, db),
	}
}

func (s *CurrencyServiceServer) GetExchangeRate(ctx context.Context, req *pb.GetExchangeRateRequest) (*pb.GetExchangeRateResponse, error) {
	rate, err := s.ctrl.GetExchangeRate(ctx, req.From, req.To)
	if err != nil {
		return nil, err
	}

	return &pb.GetExchangeRateResponse{Rate: rate}, nil
}

func (s *CurrencyServiceServer) CreateExchangeRate(ctx context.Context, req *pb.CreateExchangeRateRequest) (*pb.CreateExchangeRateResponse, error) {
	err := s.ctrl.CreateExchangeRate(ctx, req.From, req.To, req.Rate)
	if err != nil {
		return &pb.CreateExchangeRateResponse{}, err
	}

	_, err = s.ctrl.GetExchangeRateDirect(ctx, req.To, req.From)
	if err == nil {
		return &pb.CreateExchangeRateResponse{}, nil
	}

	s.log.Info("Reverse rate is not set yet, setting automatically", zap.String("from", req.To.String()), zap.String("to", req.From.String()))
	err = s.ctrl.CreateExchangeRate(ctx, req.To, req.From, 1/req.Rate)
	if err != nil {
		s.log.Warn("Couldn't automatically create reverse Exchange rate", zap.Error(err))
	}

	return &pb.CreateExchangeRateResponse{}, nil
}

func (s *CurrencyServiceServer) UpdateExchangeRate(ctx context.Context, req *pb.UpdateExchangeRateRequest) (*pb.UpdateExchangeRateResponse, error) {
	err := s.ctrl.UpdateExchangeRate(ctx, req.From, req.To, req.Rate)
	return &pb.UpdateExchangeRateResponse{}, err
}

func (s *CurrencyServiceServer) DeleteExchangeRate(ctx context.Context, req *pb.DeleteExchangeRateRequest) (*pb.DeleteExchangeRateResponse, error) {
	err := s.ctrl.DeleteExchangeRate(ctx, req.From, req.To)
	return &pb.DeleteExchangeRateResponse{}, err
}

func (s *CurrencyServiceServer) Convert(ctx context.Context, req *pb.ConversionRequest) (*pb.ConversionResponse, error) {
	amount, err := s.ctrl.Convert(ctx, req.From, req.To, req.Amount)
	if err != nil {
		return nil, err
	}

	return &pb.ConversionResponse{Amount: amount}, nil
}

func (s *CurrencyServiceServer) GetCurrencies(ctx context.Context, req *pb.GetCurrenciesRequest) (*pb.GetCurrenciesResponse, error) {
	currencies, err := s.ctrl.GetCurrencies(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.GetCurrenciesResponse{Currencies: currencies}, nil
}

func (s *CurrencyServiceServer) GetExchangeRates(ctx context.Context, req *pb.GetExchangeRatesRequest) (*pb.GetExchangeRatesResponse, error) {
	rates, err := s.ctrl.GetExchangeRates(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.GetExchangeRatesResponse{Rates: rates}, nil
}
