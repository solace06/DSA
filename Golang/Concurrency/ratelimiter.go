package main
import (
    "fmt"
    "math"
    "time"
    )

//Rate limiter using token bucket algorithm

//declaring the 2 components of the algorithm
const (
    MAX_BUCKET_SIZE float64 = 3
    REFILL_RATE int = 1
    )
    
//Creating the token bucket
type TokenBucket struct{
    currBucketSize float64
    lastRefillTimeStamp int
}

//Defining the function associated with the bucket to allow requests
func (tb *TokenBucket) allowRequest(tokens float64) bool{
    //refill the bucket
    tb.refill()
    
    //check if the bucket has enough tokens
    if tb.currBucketSize>=tokens{
        tb.currBucketSize-=tokens
        return true
    }
    
    //bucket does not enough tokens so the request is dropped
    return false
}

//Defining the function associated with the bucket to fill tokens into the bucket
func (tb *TokenBucket) refill(){
    //current time in nanoseconds
    nowTime := time.Now().Nanosecond()
    
    //calculate the number of tokens to be added in the bucket
    tokensToAdd:=((nowTime-tb.lastRefillTimeStamp)/1e9)*REFILL_RATE
    
    //refill the bucket making sure the number of tokens is less than
    //equal to the maximum size of the bucket
    tb.currBucketSize = math.Min(float64(tokensToAdd)+tb.currBucketSize,MAX_BUCKET_SIZE)
    
    //upadte the last refill time stamp
    tb.lastRefillTimeStamp=nowTime
}

func ratelimiter() {
  obj := TokenBucket{
  currBucketSize:   3,
  lastRefillTimeStamp: 0,
 }

 fmt.Printf("Request processed: %v\n", obj.allowRequest(1)) //true
 fmt.Printf("Request processed: %v\n", obj.allowRequest(1)) //true
 fmt.Printf("Request processed: %v\n", obj.allowRequest(1)) //true
 fmt.Printf("Request processed: %v\n", obj.allowRequest(1)) //false, request
}