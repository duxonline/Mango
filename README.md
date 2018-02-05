MangoCommerce is an easy to use and feature rich open source eCommerce System written in Go language. 

# Feature
## Feature 1
### Featrue 1.1

**This is a bold text**


## How to Set up AutoPostBets Tool

### Set up fiddler

1. From Fiddler menu, choose Tools -> options -> HTTPS, choose Decrypt HTTPS traffic. You wil need to confirm and install Fiddler fake certificate.

2. Create a folder c:\Fiddler to store the generated files for the tool.

3. From Fiddler menu, choose Rules -> Customize Rules and then paste the following script inside a method called OnBeforeResponse:

```
		/* Log Ladbrokes request and response here */
		var timeMS = (new Date()).getTime();

		var requestFileName = "C:\\Fiddler\\" + timeMS + "_Request.txt";
		var requestBodyFileName = "C:\\Fiddler\\" + timeMS + "_RequestBody.txt";
		
		var responseFileName = "C:\\Fiddler\\" + timeMS + "_Response.txt";
		var responseBodyFileName = "C:\\Fiddler\\" + timeMS + "_ResponseBody.txt";
		

		if( oSession.host == "lb-yang.devlb.net" && 
			oSession.uriContains("/api/1/betting/betslipBet") && 
			oSession.RequestMethod == "POST") { 
			
			oSession.utilDecodeResponse(); 
			
			oSession.SaveRequest(requestFileName ,true); 
			oSession.SaveRequestBody(requestBodyFileName); 
			
			oSession.SaveResponse(responseFileName, true);
			oSession.SaveResponseBody(responseBodyFileName);
		}
		/* End logging */

```
