export const getIngredientsForWeekAndDc = (week, dc) =>
    // mock data
    Promise.resolve({
        "data": [
            {
                "summaryId": 753810,
                "productId": "15139",
                "productName": "Tomato, Heirloom Grape 10oz",
                "productSku": "PHF-10-10594-4",
                "agreedPrice": 22,
                "package": {
                    "size": 12,
                    "uom": "unit"
                },
                "percentage": {
                    "value": 95.36
                },
                "minPercentage": {
                    "value": 95.36
                },
                "buffer": {
                    "value": 7
                },
                "week": "2017-W11",
                "yearWeekNumber": "1711",
                "qtyNeeded": "19770",
                "unitType": "CASE_TYPE",
                "unitTypeValue": 1,
                "categoryId": "87",
                "categoryName": "Produce, Herbs, Fruit",
                "supplier": {
                    "name": "Lakeside Produce",
                    "id": "685",
                    "score": {
                        "value": 0
                    },
                    "code": {
                        "value": "10164"
                    },
                    "shippingName": "Lakeside Produce"
                },
                "dcId": "12",
                "dcBobCode": "NJ",
                "orderItems": [
                    {
                        "qty": 9432,
                        "price": 22,
                        "order": {
                            "id": "46183",
                            "number": "1711NJ260672",
                            "isPreOrder": "",
                            "isSent": "1",
                            "sentAt": "2017-03-01 11:41:27",
                            "createdAt": "2017-03-01 11:36:22",
                            "supplierName": "Lakeside Produce",
                            "supplierId": "685"
                        },
                        "summaryPercentage": {
                            "value": 46.72819862718446
                        },
                        "case": {
                            "size": 12,
                            "uom": "unit"
                        },
                        "deliveryDate": "2017-03-09 06:00:00",
                        "deliveryEndDate": "2017-03-01 18:00:00",
                        "buffer": {
                            "value": 7
                        },
                        "rawQty": 9426.18
                    }
                ],
                "deliveryDates": [
                    {
                        "endDateHour": "14:00",
                        "date": "Jun 30, 2017, 10:00 AM",
                        "endDate": "Jun 5, 2017, 2:00 PM",
                        "percentage": "6.5334946566567",
                        "id": "296297"
                    }
                ],
                "purchaseItems": [
                    {
                        "qty": {
                            "value": 1317.96
                        },
                        "deliveryDate": {
                            "endDateHour": "14:00",
                            "date": "Jun 30, 2017, 10:00 AM",
                            "endDate": "Jun 5, 2017, 2:00 PM",
                            "percentage": "6.5334946566567",
                            "id": "296297"
                        },
                        "status": "INITIAL",
                        "isPreorder": 0
                    }
                ],
                "packingSize": 10,
                "isIncomplete": 1,
                "unitOfMeasure": "ounce",
                "isChanged": 0,
                "isValid": 1,
                "isFromForecast": 1,
                "summaryClass": " hasdates hasdeletedorder validated",
                "percentageSlider": "",
                "productViewLogic": {
                    "yearWeekNumber": "1711",
                    "dcId": "12",
                    "hasIncompleteOrders": 1,
                    "hasChangedTag": 0,
                    "isCompleted": 0,
                    "completedOrderItems": [
                        {
                            "qty": 2146,
                            "price": 116.15,
                            "order": {
                                "id": "46174",
                                "number": "1711NJ980302",
                                "isPreOrder": "",
                                "isSent": "1",
                                "sentAt": "2017-03-01 11:42:10",
                                "createdAt": "2017-03-01 11:36:22",
                                "supplierName": "Plainfield Fruit & Produce",
                                "supplierId": "458"
                            },
                            "summaryPercentage": {
                                "value": 50.130236032126476
                            },
                            "case": {
                                "size": 37,
                                "uom": "unit"
                            },
                            "deliveryDate": "2017-03-09 05:00:00",
                            "deliveryEndDate": "2017-03-01 06:00:00",
                            "buffer": {
                                "value": 7
                            },
                            "rawQty": 2120.9
                        }
                    ],
                    "allTags": [
                        {
                            "name": "FORECAST",
                            "type": "tag.source",
                            "color": "primary",
                            "id": "15344024"
                        }
                    ],
                    "allOrderItems": [
                        {
                            "qty": 2146,
                            "price": 116.15,
                            "order": {
                                "id": "46174",
                                "number": "1711NJ980302",
                                "isPreOrder": "",
                                "isSent": "1",
                                "sentAt": "2017-03-01 11:42:10",
                                "createdAt": "2017-03-01 11:36:22",
                                "supplierName": "Plainfield Fruit & Produce",
                                "supplierId": "458"
                            },
                            "summaryPercentage": {
                                "value": 50.130236032126476
                            },
                            "case": {
                                "size": 37,
                                "uom": "unit"
                            },
                            "deliveryDate": "2017-03-09 05:00:00",
                            "deliveryEndDate": "2017-03-01 06:00:00",
                            "buffer": {
                                "value": 7
                            },
                            "rawQty": 2120.9
                        }
                    ],
                    "totalPercentage": "115.40737169033",
                    "showOrderCheckbox": 0,
                    "unsentOrdersCount": 1,
                    "sentOrderItems": [
                        {
                            "qty": 2146,
                            "price": 116.15,
                            "order": {
                                "id": "46174",
                                "number": "1711NJ980302",
                                "isPreOrder": "",
                                "isSent": "1",
                                "sentAt": "2017-03-01 11:42:10",
                                "createdAt": "2017-03-01 11:36:22",
                                "supplierName": "Plainfield Fruit & Produce",
                                "supplierId": "458"
                            },
                            "summaryPercentage": {
                                "value": 50.130236032126476
                            },
                            "case": {
                                "size": 37,
                                "uom": "unit"
                            },
                            "deliveryDate": "2017-03-09 05:00:00",
                            "deliveryEndDate": "2017-03-01 06:00:00",
                            "buffer": {
                                "value": 7
                            },
                            "rawQty": 2120.9
                        }
                    ],
                    "unsentOrderItems": [
                        {
                            "qty": 1320,
                            "price": 22,
                            "order": {
                                "id": "63948",
                                "number": "1711NJ763558",
                                "isPreOrder": "",
                                "isSent": "",
                                "sentAt": "2017-06-05 15:01:05",
                                "createdAt": "2017-06-05 15:00:45",
                                "supplierName": "Lakeside Produce",
                                "supplierId": "685"
                            },
                            "summaryPercentage": {
                                "value": 6.533494656656677
                            },
                            "case": {
                                "size": 12,
                                "uom": "unit"
                            },
                            "deliveryDate": "2017-06-30 10:00:00",
                            "deliveryEndDate": "2017-06-05 14:00:00",
                            "buffer": {
                                "value": 7
                            },
                            "rawQty": 1317.96
                        }
                    ],
                    "summariesReadyToPurchase": [
                        "751030",
                        "753810"
                    ]
                }
            }
        ]
    })
