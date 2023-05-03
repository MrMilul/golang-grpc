package service_test

import (
	"testing"
	"context"

	"example.com/laptop_store/sample"
	"example.com/laptop_store/service"
	"example.com/laptop_store/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/stretchr/testify/require"

)
func TestServerCreateLaptop(t *testing.T){

	laptopNoID := sample.NewLaptop()
	laptopNoID.Id = ""

	laptopInvalidID := sample.NewLaptop()
	laptopInvalidID.Id = "invalid_id"

	laptopDuplicateID := sample.NewLaptop()
	storeDuplicateID := service.NewInMemoryLaptopStore()
	err := storeDuplicateID.Save(laptopDuplicateID)
	require.Nil(t, err)

	testCases := []struct{
		name string
		laptop *pb.Laptop
		store service.LaptopStore
		code codes.Code
	}{
		{
			name: "Success_with_id",
			laptop: sample.NewLaptop(),
			store: service.NewInMemoryLaptopStore(),
			code: codes.OK,
		},
		{
			name: "Success_no_id",
			laptop: laptopNoID,
			store: service.NewInMemoryLaptopStore(),
			code: codes.OK,
		},
		{
			name: "Failure_invalid_id",
			laptop: laptopInvalidID,
			store: service.NewInMemoryLaptopStore(),
			code: codes.InvalidArgument,
		},
		{
			name: "Failure_duplicate_id",
			laptop: laptopDuplicateID,
			store: storeDuplicateID,
			code: codes.AlreadyExists,
		},
	}
	

	for i := range testCases{
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T){
			req := &pb.CreateLaptopRequest{
				Laptop: tc.laptop,
			}

			service := service.NewLaptopServer(tc.store)
			res, err := service.CreateLaptop(context.Background(),req)

			if tc.code == codes.OK{
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
				if len(tc.laptop.Id) > 0{
					require.Equal(t, tc.laptop.Id, res.Id)
				}
			}else{
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.code, st.Code())
			}



		})
	}

}